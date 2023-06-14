package discovery

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type Register struct {
	EtcdAddrs   []string // etcd服务地址
	DialTimeout int      // 连接响应时间

	closeCh     chan struct{}                           // 关闭连接
	leasesID    clientv3.LeaseID                        // 租约ID
	keepAliveCh <-chan *clientv3.LeaseKeepAliveResponse // 保活

	srvInfo Server           // 服务信息
	srvTTL  int64            // 服务存活时间
	cli     *clientv3.Client // etcd连接客户端
	logger  *logrus.Logger   // 日志
}

// NewRegister 基于ETCD创建一个register
func NewRegister(etcdAddrs []string, logger *logrus.Logger) *Register {
	return &Register{
		EtcdAddrs:   etcdAddrs,
		DialTimeout: 3,
		logger:      logger,
	}
}

// Register 创建Register实例
func (r *Register) Register(srvInfo Server, ttl int64) (chan<- struct{}, error) {
	var err error
	// 分割服务IP
	if strings.Split(srvInfo.Addr, ":")[0] == "" {
		return nil, errors.New("invalid ip address")
	}

	// 初始化etcd连接客户端实例
	if r.cli, err = clientv3.New(clientv3.Config{
		Endpoints:   r.EtcdAddrs,                                // 服务节点
		DialTimeout: time.Duration(r.DialTimeout) * time.Second, // 响应时间
	}); err != nil {
		return nil, err
	}
	// 服务信息
	r.srvInfo = srvInfo
	r.srvTTL = ttl

	// 注册etcd服务
	if err = r.register(); err != nil {
		return nil, err
	}

	// 关闭连接channel
	r.closeCh = make(chan struct{})

	// 保活
	go r.keepAlive()
	return r.closeCh, nil
}

// register 创建ETCD自带的实例
func (r *Register) register() error {
	// 设定最大响应时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(r.DialTimeout)*time.Second)
	defer cancel()

	// grant申请分配租约
	leaseResp, err := r.cli.Grant(ctx, r.srvTTL)
	if err != nil {
		return err
	}
	// 记录租约ID
	r.leasesID = leaseResp.ID
	// 通过keepAlive保持租约活性，KeepAlive每500毫秒执行一次lease stream的发送，接收到发送信息回执后，更新租约，服务处于活动状态
	if r.keepAliveCh, err = r.cli.KeepAlive(context.Background(), r.leasesID); err != nil {
		return err
	}
	// 序列化服务信息
	data, err := json.Marshal(r.srvInfo)
	if err != nil {
		return err
	}
	// 将服务信息以k-v方式存储
	_, err = r.cli.Put(context.Background(), BuildRegisterPath(r.srvInfo), string(data), clientv3.WithLease(r.leasesID))
	if err != nil {
		return err
	}
	return err
}

// keepAlive 保持租约
func (r *Register) keepAlive() {
	ticker := time.NewTicker(time.Duration(r.srvTTL) * time.Second)
	for {
		select {
		case <-r.closeCh:
			// 删除节点
			if err := r.unregister(); err != nil {
				fmt.Println("unregister failed error")
			}
			// 撤销租约，所有附加到租约里的key将过期并删除
			if _, err := r.cli.Revoke(context.Background(), r.leasesID); err != nil {
				fmt.Println("revoke fail")
			}
		case res := <-r.keepAliveCh:
			if res == nil {
				if err := r.register(); err != nil {
					fmt.Println("register err")
				}
			}
		case <-ticker.C:
			if r.keepAliveCh == nil {
				if err := r.register(); err != nil {
					fmt.Println("register err")
				}
			}

		}
	}
}

// 删除节点
func (r *Register) unregister() error {
	_, err := r.cli.Delete(context.Background(), BuildRegisterPath(r.srvInfo))
	return err
}
