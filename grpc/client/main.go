package main

import (
	"context"
	"log"
	"time"

	pb "auth-grpc-golang/grpc/proto"

	"google.golang.org/grpc"
)

const (
	address     = "localhost:50051"
	defaultName = "Connected"
)

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewAddServiceClient(conn)

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SinUp(ctx, &pb.SinUpRequest{EmailId: "Ebrahim"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Connect: %s", r.GetMessage())

	s, err := c.SinIn(ctx, &pb.SinInRequest{EmailId: "ebrahims1902@gmail.com"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Connect: %s", s.GetMessage())

}
