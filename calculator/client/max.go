package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.CalculatorServiceClient) {
	log.Println("doMax func was invoked")

	stream, err := c.Max(context.Background())

	if err != nil {
		log.Fatalf("error while opening stream: %v \n", err)
	}
	waitChannel := make(chan struct{})

	go func() {
		numbers := []int32{1, 8, 7, 9, 15, 88, 10}
		for _, num := range numbers {
			log.Printf("sending number: %d \n", num)
			stream.Send(&pb.MaxRequest{
				Number: num,
			})
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
				log.Fatalf("error while reading server streamL %v \n", err)
			}
			log.Printf("reveived a new maximum: %d \n", res.Result)

		}
		close(waitChannel)
	}()

	<-waitChannel
}
