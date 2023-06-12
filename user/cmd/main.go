package main

import (
	"net"
	"user/config"
	"user/discovery"
	"user/internal/handler"
	"user/internal/repository"
	"user/internal/service"

	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

func main() {
	config.InitConfig() // 初始化配置文件
	repository.InitDB() // 初始化数据库

	// etcd地址
	etcdAddress := []string{viper.GetString("etcd.address")}
	// 服务注册
	etcdRegister := discovery.NewRegister(etcdAddress, logrus.New())
	grpcAddress := viper.GetString("server.grpcAddress")
	userNode := discovery.Server{
		Name: viper.GetString("server.domain"),
		Addr: grpcAddress,
	}
	server := grpc.NewServer()
	defer server.Stop()

	// 绑定服务
	service.RegisterUserServiceServer(server, handler.NewUserService())
	listener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		panic(err)
	}
	if _, err = etcdRegister.Register(userNode, 10); err != nil {
		panic(err)
	}
	if err = server.Serve(listener); err != nil {
		panic(err)
	}

}
