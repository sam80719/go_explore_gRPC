package main

import (
	pb "github.com/sam80719/go_explore_gRPC/greeting/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	// start add ssl cerf
	opts := []grpc.ServerOption{}
	tls := true

	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"

		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)

		if err != nil {
			log.Fatalf("failed loading cerficates: %v \n", err)
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)
	pb.RegisterGreetingServiceServer(s, &Service{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v \n", err)
		return
	}
}
