package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/greeting/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetingWithDeadline(c pb.GreetingServiceClient, timeout time.Duration) {
	log.Println("doGreetingWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()

	req := &pb.GreetingRequest{FirstName: "Sam"}

	res, err := c.GreetingWithDeadLine(ctx, req)

	if err != nil {
		e, ok := status.FromError(err)

		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("deadline exceeded")
				return
			} else {
				log.Fatalf("unexpected gRPC errorL %v \n", err)
			}
		} else {
			log.Fatalf("a non gRPC error: %v \n", err)
		}
	}
	log.Printf("greetingWithDeadline: %s \n", res.Result)
}
