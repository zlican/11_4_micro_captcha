package main

import (
	"user/handler"
	"user/model"
	pb "user/proto"

	"github.com/go-micro/plugins/v4/registry/consul"
	"go-micro.dev/v4"
	"go-micro.dev/v4/logger"
)

var (
	service = "user"
	version = "latest"
)

func main() {
	//初始化 redisPool
	model.InitRedis()

	consulReg := consul.NewRegistry()

	// Create service
	srv := micro.NewService(
		micro.Address("127.0.0.1:12342"),
		micro.Registry(consulReg),
	)
	srv.Init(
		micro.Name(service),
		micro.Version(version),
	)

	// Register handler
	if err := pb.RegisterUserHandler(srv.Server(), new(handler.User)); err != nil {
		logger.Fatal(err)
	}
	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
