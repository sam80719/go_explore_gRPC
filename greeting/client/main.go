package main

import (
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
}
