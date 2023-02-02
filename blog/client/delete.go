package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/blog/proto"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Println("--deleteBlog was invoked")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})

	if err != nil {
		log.Fatalf("error while deleting: %v \n", err)
	}

	log.Println("blog was deleted!")

}
