package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

var collection *mongo.Collection
var address string = "0.0.0.0:50052"

type Service struct {
	pb.BlogServiceServer
}

func main() {
	// set up mongo url
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatal(err)
		return
	}

	// set mongo connect
	err = client.Connect(context.Background())
	if err != nil {
		log.Fatal(err)
		return
	}

	// select db and collection
	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("failed to listen on: %v \n", err)
		return
	}
	log.Printf("listening on %s \n", address)

	s := grpc.NewServer()
	pb.RegisterBlogServiceServer(s, &Service{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to server: %v \n", err)
		return
	}
}
