package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/calculator/proto"
	"log"
)

func doSum(c pb.CalculatorServiceClient) {
	log.Println("doSum was invoked")
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  2,
		SecondNumber: 9,
	})

	if err != nil {
		log.Fatalf("could not sum: %v", err)
	}

	log.Printf("SumL %d \n", res.Result)

}
