package main

import (
	"api-gateway/config"
	"api-gateway/discovery"
	"api-gateway/internal/service"
	"api-gateway/routes"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"google.golang.org/grpc"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/resolver"

	"github.com/spf13/viper"
)

func main() {
	config.InitConfig()
	// 服务发现
	etcdAddress := []string{viper.GetString("etcd.address")}
	fmt.Println("etcdAddress: ", etcdAddress)
	etcdRegister := discovery.NewResolver(etcdAddress, logrus.New())
	resolver.Register(etcdRegister)
	go startListen()
	{
		osSignal := make(chan os.Signal, 1)
		signal.Notify(osSignal, os.Interrupt, os.Kill, syscall.SIGTERM, syscall.SIGINT, syscall.SIGKILL)
		s := <-osSignal
		fmt.Println("exit", s)
	}
	fmt.Println("gateway listen on:4000")

}

func startListen() {
	opts := []grpc.DialOption{
		grpc.WithInsecure(),
	}
	// grpc远程连接user模块端口
	userGrpcAddress := viper.GetString("user.grpcAddress")
	userConn, err := grpc.Dial(userGrpcAddress, opts...)
	if err != nil {
		panic(err)
	}
	// 初始化user模块客户端连接
	userService := service.NewUserServiceClient(userConn)

	// grpc远程连接task模块端口
	taskGrpcAddress := viper.GetString("task.grpcAddress")
	taskConn, err2 := grpc.Dial(taskGrpcAddress, opts...)
	if err2 != nil {
		panic(err2)
	}
	// 初始化user模块客户端连接
	taskService := service.NewTaskServiceClient(taskConn)

	ginRouter := routes.NewRouter(userService, taskService)
	server := &http.Server{
		Addr:           viper.GetString("server.port"),
		Handler:        ginRouter,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	err3 := server.ListenAndServe()
	if err3 != nil {
		fmt.Println("绑定失败，可能端口被占用", err3)
		panic(err2)
	}
}
