package main

import (
	pb "github.com/sam80719/go_explore_gRPC/greeting/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var address string = "0.0.0.0:50052"

type Service struct {
	pb.GreetingServiceServer
}

func main() {
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen on: %v \n", err)
		return
	}
	log.Printf("listening on %s \n", address)

	s := grpc.NewServer()

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v \n", err)
		return
	}
}
