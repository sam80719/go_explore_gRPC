package main

import (
	"context"
	"fmt"
	pb "github.com/sam80719/go_explore_gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Service) ListBlogs(in *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {
	log.Println("listBlogs was invoked")

	current, err := collection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknow internal error: %v", err),
		)
	}

	defer current.Close(context.Background())

	for current.Next(context.Background()) {
		data := &BlogItem{}
		err := current.Decode(data)

		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("error while decoding data from mongoDB: %v \n", err),
			)
		}
		stream.Send(documentToBlog(data))
	}

	if err = current.Err(); err != nil {
		return status.Errorf(
			codes.Internal,
			fmt.Sprintf("unknow internal error: %v \n", err),
		)
	}
	return nil
}
