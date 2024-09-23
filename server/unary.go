package main

import (
	"context"
	"log"

	pb "github.com/tejasvi541/Go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	log.Printf("Received: %v", req)
	return &pb.HelloResponse{Message: "Hello from the server!"}, nil
}

