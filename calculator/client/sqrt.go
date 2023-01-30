package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func doSqrt(c pb.CalculatorServiceClient, n int32) {
	log.Println("doSqrt func was invoked!")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{Number: n})

	// handle err
	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			log.Printf("error msg from server: %s \n", e.Message())
			log.Printf("error code from server: %s \n", e.Code())

			if e.Code() == codes.InvalidArgument {
				log.Println("it's seem to send a negative number!")
				return
			}
		} else {
			log.Fatalf("a non grpc error: %v \n", err)
		}
	}

	log.Printf("sqrt: %+v \n", res.Result)
}
