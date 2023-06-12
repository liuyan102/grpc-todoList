package main

import (
	"net"
	"task/config"
	"task/discovery"
	"task/internal/handler"
	"task/internal/repository"
	"task/internal/service"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {
	config.InitConfig() // 初始化配置文件
	repository.InitDB() // 初始化数据库

	// etcd地址
	etcdAddress := []string{viper.GetString("etcd.address")}
	// 注册etcd服务
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	// grpc地址
	grpcAddress := viper.GetString("server.grpcAddress")
	// 服务节点
	taskNode := discovery.Server{
		Name: viper.GetString("server.domain"),
		Addr: grpcAddress,
	}
	// 新建grpc服务
	server := grpc.NewServer()
	defer server.Stop()

	// 绑定服务
	service.RegisterTaskServiceServer(server, handler.NewTaskService())
	// 设置监听器监听端口
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	// 注册etcd节点
	if _, err = etcdRegister.Register(taskNode, 10); err != nil {
		panic(err)
	}
	// 开始监听端口
	if err = server.Serve(listener); err != nil {
		panic(err)
	}
}
