package config

import (
	"context"
	"log"
)

type Server struct {
	Change   chan string
	Response chan string
}

func (s *Server) ReloadConfig(ctx context.Context, message *Message) (*Message, error) {
	log.Println("Received message body from client:", message.Body)
	s.Change <- "updateconfig"
	// ticker := time.NewTicker(10 * time.Second)
	// go func() {
	// 	select {
	// 	case <-ticker.C:
	// 		fmt.Println("cannot receive the msg in 10s")
	// 		return &Message{Body: "cannot receive the msg in 10s"}, nil
	// 	}
	// }()
	res := <-s.Response
	return &Message{Body: res}, nil
}

func (s *Server) ApprovalToOneSplitAudit(ctx context.Context, message *Message) (*Message, error) {
	log.Println("Received message body from client:", message.Body)
	s.Change <- message.Body
	// ticker := time.NewTicker(10 * time.Second)
	// go func() {
	// 	select {
	// 	case <-ticker.C:
	// 		log.Fatal("cannot receive the msg in 10s")
	// 	}
	// }()
	res := <-s.Response
	return &Message{Body: res}, nil
}
