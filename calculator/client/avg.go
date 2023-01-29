package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/calculator/proto"
	"log"
)

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("doAvg func was invoked!")
	stream, err := c.Avg(context.Background())

	if err != nil {
		log.Fatalf("error while opening the stream %v \n", err)
	}

	numbers := []int32{7, 5, 12, 9, 55, 77}

	for _, number := range numbers {
		log.Printf("sending number is %d \n", number)
		stream.Send(&pb.AvgRequest{Number: number})
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response: %v \n", err)
	}

	log.Printf("avg: %f \n", res.Result)

}
