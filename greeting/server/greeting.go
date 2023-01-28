package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/greeting/proto"
	"log"
)

func (s *Service) Greeting(ctx context.Context, in *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	log.Printf("Greeting Finction was invoked with %v! \n", in)
	return &pb.GreetingResponse{
		Result: "Hello " + in.FirstName,
	}, nil
}
