package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/greeting/proto"
	"log"
	"time"
)

func doLongGreeting(c pb.GreetingServiceClient) {
	log.Println("doLongGreeting func was invoked!")

	reqs := []*pb.GreetingRequest{
		{FirstName: "Sam"},
		{FirstName: "Jimmy"},
		{FirstName: "Kevin"},
	}

	stream, err := c.LongGreeting(context.Background())

	if err != nil {
		log.Fatalf("error while calling longGreeting: %v \n", err)
	}

	for _, req := range reqs {
		log.Printf("sending req: %v \n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error while receiving response from longGreeting %v \n", err)
	}
	log.Printf("longGreeting: %s \n", res.Result)
}
