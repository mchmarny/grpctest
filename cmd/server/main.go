package main

import (
	"fmt"
	"log"
	"net"

	v1 "github.com/mchmarny/grpctest/pkg/api/v1"

	"google.golang.org/grpc"
)

// main start a gRPC server and waits for connection
func main() {
	// create a listener on TCP port 7777
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 7777))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// create a server instance
	s := v1.Server{}

	// create a gRPC server object
	grpcServer := grpc.NewServer()

	// attach the Ping service to the server
	v1.RegisterPingServer(grpcServer, &s)

	// run HTTP gateway
	go func() {
		_ = v1.RunServer(ctx, fmt.Sprintf(":%d", 7777), fmt.Sprintf(":%d", 8888))
	}()

	// start the server
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
