package main

import (
	"log"
	"time"

	pb "github.com/tejasvi541/Go-grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer)error{
	log.Printf("Got the Names List: %v", req.Names)
	for _, name := range req.Names{
		res := &pb.HelloResponse{Message: "Hello "+name}
		if err := stream.Send(res); err != nil{
			return err
		}
		time.Sleep(1*time.Second)
	}
	return nil

}