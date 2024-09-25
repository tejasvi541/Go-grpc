package main

import (
	"context"
	"log"
	"time"

	pb "github.com/tejasvi541/Go-grpc/proto"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Calling SayHelloClientStreaming")
	stream, err := client.SayHelloClientStreaming(context.Background())

	if err != nil {
		log.Fatalf("Error while calling SayHelloClientStreaming: %v", err)
	}

	for _, name := range names.Names {
		log.Printf("Sending message: %s", name)


		if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
			log.Fatalf("Error while sending message: %v", err)
		}
		time.Sleep(2000 * time.Millisecond)
	}
	res, err := stream.CloseAndRecv() 
	if err != nil {
		log.Fatalf("Error while receiving response from SayHelloClientStreaming: %v", err)
	}
	log.Printf("Response from SayHelloClientStreaming: %v", res.Messages)
}