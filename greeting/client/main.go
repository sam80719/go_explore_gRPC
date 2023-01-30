package main

import (
	pb "github.com/sam80719/go_explore_gRPC/greeting/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"log"
)

var address string = "localhost:50052"

func main() {
	// start add ssl cerf
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")

		if err != nil {
			log.Fatalf("error while loading CA trust certificate: %v \n", err)
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	coon, err := grpc.Dial(address, opts...)

	if err != nil {
		log.Fatalf("failed to connectL %v \n", err)
		return
	}
	defer coon.Close()

	c := pb.NewGreetingServiceClient(coon)
	doGreeting(c)
	//doGreetingManyTime(c)
	//doLongGreeting(c)
	//doGreetingEveryone(c)
	//doGreetingWithDeadline(c, 2*time.Second)
}
