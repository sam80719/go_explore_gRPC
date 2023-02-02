package main

import (
	pb "github.com/sam80719/go_explore_gRPC/blog/proto"
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

	c := pb.NewBlogServiceClient(coon)

	// unary
	id := createBlog(c)

	readBlog(c, id) // valid
	readBlog(c, "notExistId")

}
