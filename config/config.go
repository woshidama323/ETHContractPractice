package config

import (
	"context"
	"log"
)

type Server struct {
	Change chan bool
}

func (s *Server) ReloadConfig(ctx context.Context, message *Message) (*Message, error) {
	log.Println("Received message body from client:", message.Body)
	s.Change <- true
	return &Message{Body: "Hello from server"}, nil
}
