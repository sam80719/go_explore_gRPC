package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/greeting/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func (s *Service) GreetingWithDeadLine(ctx context.Context, in *pb.GreetingRequest) (*pb.GreetingResponse, error) {
	log.Printf("greetingWithDeadline func was invoked with %v \n", in)

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			log.Println("the client canceled the request!")
			return nil, status.Error(codes.Canceled, "the client canceled the request")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetingResponse{
		Result: "hello " + in.FirstName,
	}, nil

}
