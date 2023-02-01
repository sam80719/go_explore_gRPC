package main

import (
	"context"
	"fmt"
	pb "github.com/sam80719/go_explore_gRPC/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	fmt.Println("-- createBlog func was invoked!")

	blog := &pb.Blog{
		AuthorId: "Sam",
		Title:    "first time to create title",
		Content:  "first tome to create blog content",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("unexpected error: %v \n", err)
	}

	log.Printf("blog content has been created: %s \n", res.Id)
	return res.Id
}
