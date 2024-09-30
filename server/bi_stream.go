package main

import (
	"io"
	"log"

	pb "github.com/tejasvi541/Go-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			break
		}
		log.Printf("Received message %s", req.Name)
		res := &pb.HelloResponse{Message: "Hello " + req.Name}
		if err:=stream.Send(res); err != nil {
			log.Printf("Error sending message: %v", err)
			break
		}
	}
	return nil
}