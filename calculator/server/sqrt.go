package main

import (
	"context"
	"fmt"
	pb "github.com/sam80719/go_explore_gRPC/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
)

func (s *Service) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Println("sqrt func was invoked")

	number := in.Number

	if number < 0 {
		// export err
		return nil, status.Error(
			codes.InvalidArgument,
			fmt.Sprintf("received a negative number: %d \n", number),
		)
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(number)),
	}, nil
}
