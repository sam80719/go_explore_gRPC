package main

import (
	pb "github.com/sam80719/go_explore_gRPC/calculator/proto"
	"io"
	"log"
)

func (s *Service) Max(stream pb.CalculatorService_MaxServer) error {
	log.Println("max func was invoked!")
	var maximum int32 = 0

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("error while reading client stream: %v \n", err)
		}

		if num := req.Number; num > maximum {
			maximum = num
			err := stream.Send(&pb.MaxResponse{
				Result: maximum,
			})

			if err != nil {
				log.Fatalf("error while seding data to client %v \n", err)
			}

		}
	}

}
