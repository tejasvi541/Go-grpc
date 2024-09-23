package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tejasvi541/Go-grpc/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	defer cancel()

	res, err := client.SayHello(ctx, &pb.NoParam{})

	if err != nil {
		log.Fatalf("Failed to call SayHello: %v", err)
	}
	log.Printf("Response: %v", res.Message)
}