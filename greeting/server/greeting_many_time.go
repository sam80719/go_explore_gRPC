package main

import (
	"fmt"
	pb "github.com/sam80719/go_explore_gRPC/greeting/proto"
	"log"
)

func (s *Service) GreetingManyTime(in *pb.GreetingRequest, stream pb.GreetingService_GreetingManyTimeServer) error {
	log.Printf("GreetingManyTime func was invoked with: %v \n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("hello %s, number is %d", in.FirstName, i)

		stream.Send(&pb.GreetingResponse{
			Result: res,
		})
	}
	return nil
}
