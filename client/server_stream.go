package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tejasvi541/Go-grpc/proto"
)

func callSayHelloServerStream(client pb.GreetServiceClient, names *pb.NamesList) {
	ctx , cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	stream , err := client.SayHelloServerStreaming(ctx, names)

	if err != nil {
		log.Fatalf("Failed to call SayHelloServerStreaming: %v", err)
	}

	for {
		res, err := stream.Recv()
		if err != nil {
			log.Fatalf("Failed to receive from stream: %v", err)
		}
		log.Printf("Response: %v", res.Message)
	}
}