package discovery

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc/resolver"
)

// 此resolver支持的schema方法
const schema = "etcd"

// Resolver 解析器
type Resolver struct {
	schema      string   // 方法
	EtcdAddrs   []string // etcd服务地址
	DialTimeout int      // 响应时间

	closeCh      chan struct{}      // 关闭连接channel
	watchCh      clientv3.WatchChan // 监听channel
	cli          *clientv3.Client   // etcd连接实例
	keyPrefix    string             // 关键字
	srvAddrsList []resolver.Address // grpc服务地址

	cc     resolver.ClientConn // ClientConn包含解析器的回调函数，以通知gRPC ClientConn的任何更新。
	logger *logrus.Logger      // 日志
}

// NewResolver 新建resolver实例
func NewResolver(etcdAddress []string, logger *logrus.Logger) *Resolver {
	return &Resolver{
		schema:      schema,
		EtcdAddrs:   etcdAddress,
		DialTimeout: 3,
		logger:      logger,
	}
}

// Scheme 返回此resolver支持的方法.
func (r *Resolver) Scheme() string {
	return r.schema
}

// Build 给指定的target创建一个新的resolver
func (r *Resolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r.cc = cc
	r.keyPrefix = BuildPrefix(Server{Name: target.Endpoint(), Version: target.Authority})
	if _, err := r.start(); err != nil {
		return nil, err
	}
	return r, nil
}

// ResolveNow 被 gRPC 调用，以尝试再次解析目标名称。只用于提示，可忽略该方法;
func (r *Resolver) ResolveNow(o resolver.ResolveNowOptions) {}

// Close resolver.Resolver interface
func (r *Resolver) Close() {
	r.closeCh <- struct{}{}
}

// start
func (r *Resolver) start() (chan<- struct{}, error) {
	var err error
	// 新建etcd连接实例
	r.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.EtcdAddrs,
		DialTimeout: time.Duration(r.DialTimeout) * time.Second,
	})
	if err != nil {
		return nil, err
	}
	// 注册解析器
	resolver.Register(r)
	r.closeCh = make(chan struct{})
	// 同步获取的所有地址
	if err = r.sync(); err != nil {
		return nil, err
	}
	// 监听
	go r.watch()

	return r.closeCh, nil
}

// watch 监听前缀的信息变更，有变更的通知，及时更新srvAddrsList中的地址信息
func (r *Resolver) watch() {
	ticker := time.NewTicker(time.Minute)
	r.watchCh = r.cli.Watch(context.Background(), r.keyPrefix, clientv3.WithPrefix())

	for {
		select {
		case <-r.closeCh:
			return
		case res, ok := <-r.watchCh:
			if ok {
				r.update(res.Events)
			}
		case <-ticker.C:
			if err := r.sync(); err != nil {
				r.logger.Error("sync failed", err)
			}
		}
	}
}

// update 更新srvAddrsList中的地址信息
func (r *Resolver) update(events []*clientv3.Event) {
	for _, ev := range events {
		var info Server
		var err error

		switch ev.Type {
		// 新增地址信息
		case clientv3.EventTypePut:
			info, err = ParseValue(ev.Kv.Value)
			if err != nil {
				continue
			}
			addr := resolver.Address{Addr: info.Addr, Metadata: info.Weight}
			if !Exist(r.srvAddrsList, addr) {
				r.srvAddrsList = append(r.srvAddrsList, addr)
				r.cc.UpdateState(resolver.State{Addresses: r.srvAddrsList})
			}
		// 删除地址信息
		case clientv3.EventTypeDelete:
			info, err = SplitPath(string(ev.Kv.Key))
			if err != nil {
				continue
			}
			addr := resolver.Address{Addr: info.Addr}
			if s, ok := Remove(r.srvAddrsList, addr); ok {
				r.srvAddrsList = s
				r.cc.UpdateState(resolver.State{Addresses: r.srvAddrsList})
			}
		}
	}
}

// sync 同步获取所有地址信息
func (r *Resolver) sync() error {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	res, err := r.cli.Get(ctx, r.keyPrefix, clientv3.WithPrefix())
	if err != nil {
		return err
	}
	r.srvAddrsList = []resolver.Address{}

	// 定时同步etcd中可用的服务地址到srvAddrsList中
	for _, v := range res.Kvs {
		info, err := ParseValue(v.Value)
		if err != nil {
			continue
		}
		addr := resolver.Address{Addr: info.Addr, Metadata: info.Weight}
		r.srvAddrsList = append(r.srvAddrsList, addr)
	}
	// 更新ClientConn中的Address
	err = r.cc.UpdateState(resolver.State{Addresses: r.srvAddrsList})
	if err != nil {
		return err
	}
	return nil
}
