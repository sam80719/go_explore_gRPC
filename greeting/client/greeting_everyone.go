package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/greeting/proto"
	"io"
	"log"
	"time"
)

func doGreetingEveryone(c pb.GreetingServiceClient) {
	log.Println("doGreetingEveryone func was invoked!")

	stream, err := c.GreetingEveryone(context.Background())

	if err != nil {
		log.Fatalf("error while creating stream: %v \n", err)
	}

	reqs := []*pb.GreetingRequest{
		{FirstName: "Sam"},
		{FirstName: "Kim"},
		{FirstName: "Jimmy"},
	}

	// force the function to wait before exiting until the second go routine closes the channel
	waitChannel := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("sending request: %v \n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("error while receiving: %v \n", err)
			}
			log.Printf("received: %v \n", res.Result)
		}
		close(waitChannel)
	}()
	<-waitChannel
}
