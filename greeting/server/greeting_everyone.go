package main

import (
	pb "github.com/sam80719/go_explore_gRPC/greeting/proto"
	"io"
	"log"
)

func (s *Service) GreetingEveryone(stream pb.GreetingService_GreetingEveryoneServer) error {
	log.Println("greetingEveryone func was invoked!!")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading client stream %v! \n", err)
		}

		res := "hallo " + req.FirstName + "!"
		err = stream.Send(&pb.GreetingResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("error while sending data to client: %v \n", err)
		}

	}
}
