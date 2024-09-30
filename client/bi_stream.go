package main

import (
	"context"
	"io"
	"log"

	pb "github.com/tejasvi541/Go-grpc/proto"
)

func callSayHelloBidirectionalStreaming(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Println("Calling SayHelloBidirectionalStreaming")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())

	if err != nil {
		log.Fatalf("Error while calling SayHelloBidirectionalStreaming: %v", err)
	}

	waitc := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				close(waitc)
				return
			}
			if err != nil {
				log.Fatalf("Error while receiving message: %v", err)
				return
			}
			log.Printf("Received message: %s", res.Message)
		}
	}()
	go func() {
		for _, name := range names.Names {
			log.Printf("Sending message: %s", name)
			if err := stream.Send(&pb.HelloRequest{Name: name}); err != nil {
				log.Fatalf("Error while sending message: %v", err)
			}
		}
		stream.CloseSend()
	}()
	<-waitc
}