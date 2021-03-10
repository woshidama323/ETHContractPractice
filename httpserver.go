package main

import (
	"net"

	"github.com/woshidama323/config"
	"google.golang.org/grpc"
)

//GrpcServer 服务初始化
func GrpcServer(change, response chan string) {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		rlog.Fatalln("failed to listen for grpc server,err:", err)
	}

	grpcServer := grpc.NewServer()

	s := config.Server{
		Change:   change,
		Response: response,
	}
	config.RegisterConfigUpdateServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		rlog.Fatalln("failed to serve: %s", err)
	}

}
