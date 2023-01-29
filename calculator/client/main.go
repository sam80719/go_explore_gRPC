package main

import (
	pb "github.com/sam80719/go_explore_gRPC/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var address string = "localhost:50052"

func main() {
	// transport security: grpc.WithTransportCredentials(insecure.NewCredentials())
	coon, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("failed to connectL %v \n", err)
		return
	}
	defer coon.Close()

	c := pb.NewCalculatorServiceClient(coon)

	// unary
	//doSum(c)

	// server stream
	//doPrimes(c)

	// client stream
	//doAvg(c)

	// Bi-directional Streaming
	doMax(c)
}
