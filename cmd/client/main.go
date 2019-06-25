package main

import (
	"flag"
	"log"

	v1 "github.com/mchmarny/grpctest/pkg/api/v1"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

var (
	uri  = flag.String("server", ":8080", "Server URI (eg app.domain.com:8080")
	name = flag.String("name", "foo", "Name to say hi to")
)

func main() {

	flag.Parse()

	ctx := context.Background()

	log.Printf("uri: %s", *uri)
	log.Printf("name: %s", *name)

	var conn *grpc.ClientConn
	conn, err := grpc.Dial(*uri, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %s", err)
	}
	defer conn.Close()

	msg := &v1.PingMessage{
		Greeting: *name,
	}

	c := v1.NewPingClient(conn)
	response, err := c.Say(ctx, msg)
	if err != nil {
		log.Fatalf("Error when calling Say: %s", err)
	}

	log.Printf("Response: %s", response.Greeting)
}
