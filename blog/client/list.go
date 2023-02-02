package main

import (
	"context"
	pb "github.com/sam80719/go_explore_gRPC/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("--listBlog func was invoked")

	stream, err := c.ListBlogs(context.Background(), &emptypb.Empty{})

	if err != nil {
		log.Fatalf("error whilr calling listBlog: %v \n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("something happend: %v \n", err)
		}

		log.Println(res)
	}
}
