package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/calculator/proto"
	"log"
)

func (s *Service) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Sum function was invoked with %v \n", in)
	return &pb.SumResponse{
		Result: in.FirstNumber + in.SecondNumber,
	}, nil
}
