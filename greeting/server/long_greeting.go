package main

import (
	"fmt"
	pb "github.com/sam80719/go_explore_gRPC/greeting/proto"
	"io"
	"log"
)

func (s *Service) LongGreeting(stream pb.GreetingService_LongGreetingServer) error {
	log.Println("longGreeting func was invoked!")

	res := ""

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetingResponse{Result: res})
		}

		if err != nil {
			log.Fatalf("error while reading client stream: %v \n", err)
		}
		log.Printf("Receiving: %v \n", req)
		res += fmt.Sprintf("hallo %s \n", req.FirstName)
	}
}
