package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/greeting/proto"
	"io"
	"log"
)

func doGreetingManyTime(c pb.GreetingServiceClient) {
	log.Println("doGreetingManyTime was invoked!")
	req := &pb.GreetingRequest{FirstName: "Sam"}

	stream, err := c.GreetingManyTime(context.Background(), req)

	if err != nil {
		log.Fatalf("error while calling GreetingManyTime: %v \n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("error while reading stream: %v \n", err)
		}

		log.Printf("GreetingManyTime: %s \n", msg.Result)
	}
}
