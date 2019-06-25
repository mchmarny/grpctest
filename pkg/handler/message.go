package handler

import (
	"log"

	"golang.org/x/net/context"

	v1 "github.com/mchmarny/grpctest/pkg/api/v1"
)

// Server represents the gRPC server
type Server struct{}

// Say generates response to a Ping request
func (s *Server) Say(ctx context.Context, in *v1.PingMessage) (*v1.PingMessage, error) {

	log.Printf("Receive message: %s", in.Greeting)

	return &v1.PingMessage{
		Greeting: "Hi " + in.Greeting,
	}, nil

}
