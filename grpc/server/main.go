package main

import (
	"auth-grpc-golang/grpc/proto"
	pb "auth-grpc-golang/grpc/proto"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedAddServiceServer
}

func (s *server) SinUp(ctx context.Context, in *pb.SinUpRequest) (*pb.SinUpResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.SinUpResponse{Message: "Success"}, nil
}

func (s *server) SinIn(ctx context.Context, in *pb.SinInRequest) (*pb.SinInResponse, error) {
	log.Println("Received:", in)
	return &pb.SinInResponse{Message: "Successfully SinIn"}, nil
}

func main() {
	listener, err := net.Listen("tcp", port)
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	proto.RegisterAddServiceServer(srv, &server{})
	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}
