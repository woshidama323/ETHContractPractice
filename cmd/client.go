package main

import (
	"context"
	"log"

	"github.com/woshidama323/config"
	"google.golang.org/grpc"
)

func main() {
	var conn *grpc.ClientConn
	conn, err := grpc.Dial(":9000", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("failed to dial server,err:", err)
	}

	defer conn.Close()

	c := config.NewConfigUpdateClient(conn)
	message := config.Message{
		Body: "hello from client !",
	}

	response, err := c.ReloadConfig(context.Background(), &message)
	if err != nil {
		log.Fatalf("Error when calling reloadconfig,err:%s", err)
	}
	log.Println("Response from server:", response)
}
