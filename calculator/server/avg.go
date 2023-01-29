package main

import (
	pb "github.com/sam80719/go_explore_gRPC/calculator/proto"
	"io"
	"log"
)

func (s *Service) Avg(stream pb.CalculatorService_AvgServer) error {
	var sum int32 = 0
	count := 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{
				Result: float64(sum) / float64(count),
			})
		}

		if err != nil {
			log.Fatalf("error while reading client stream %v \n", err)
		}
		log.Printf("receiving number is %d \n", req.Number)
		sum += req.Number
		count++
	}
}
