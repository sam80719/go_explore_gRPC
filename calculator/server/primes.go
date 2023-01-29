package main

import (
	pb "github.com/sam80719/go_explore_gRPC/calculator/proto"
	"log"
)

func (s *Service) Primes(in *pb.PrimeRequest, stream pb.CalculatorService_PrimesServer) error {
	log.Printf("primes func was invoked with %v \n", in)

	number := in.Number
	divisor := int64(2)

	for number > 1 {
		if number%divisor == 0 {
			stream.Send(&pb.PrimeResponse{Result: divisor})
			number /= divisor
		} else {
			divisor++
		}
	}
	return nil
}
