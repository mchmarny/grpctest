package main

import (
	"fmt"
	"log"
	"net"
	"os"

	v1 "github.com/mchmarny/grpctest/pkg/api/v1"

	"github.com/mchmarny/grpctest/pkg/handler"

	"google.golang.org/grpc"
)

const (
	defaultPort      = "8080"
	portVariableName = "PORT"
)

// main start a gRPC server and waits for connection
func main() {

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// create a listener on TCP port
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// create a server instance
	s := handler.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	v1.RegisterPingServer(grpcServer, &s)

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
