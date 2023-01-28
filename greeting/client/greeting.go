package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/greeting/proto"
	"log"
)

func doGreeting(c pb.GreetingServiceClient) {
	log.Println("doGreeting was invoked")
	res, err := c.Greeting(context.Background(), &pb.GreetingRequest{
		FirstName: "sam",
	})
	if err != nil {
		log.Fatalf("could not greeting: %v \n", err)
	}
	log.Printf("Greeting: %s \n", res.Result)
}
