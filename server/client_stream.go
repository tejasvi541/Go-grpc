package main

import (
	"io"
	"log"

	pb "github.com/tejasvi541/Go-grpc/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for{
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessagesList{Messages: messages})
		}
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			break
		}
		log.Printf("Received message %s", req.Name)
		messages = append(messages, "Heuyyyoo" , req.Name)
	}	
	return nil
}