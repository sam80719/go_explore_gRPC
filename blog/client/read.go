package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("--readBlog function was invoked")

	req := &pb.BlogId{Id: id}

	res, err := c.ReadBlog(context.Background(), req)

	if err != nil {
		log.Fatalf("error while reading: %v \n", err)
	}
	log.Printf("blog was read: %v \n", res)
	return res

}
