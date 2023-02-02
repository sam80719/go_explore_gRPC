package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("--updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "new sam",
		Title:    "a new title",
		Content:  "add new awesome condition!!!!!!!",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)

	if err != nil {
		log.Fatalf("error while updating: %v \n", err)
	}

	log.Println("blog was updated!!")
}
