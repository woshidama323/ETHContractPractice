package main

import (
	"log"
	"net"

	"github.com/woshidama323/config"
	"google.golang.org/grpc"
)

//GrpcServer 服务初始化
func GrpcServer(change chan bool) {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen for grpc server,err:", err)
	}

	grpcServer := grpc.NewServer()

	s := config.Server{
		Change: change,
	}
	config.RegisterConfigUpdateServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}

}
