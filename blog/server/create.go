package main

import (
	"context"
	"fmt"
	pb "github.com/sam80719/go_explore_gRPC/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Service) CreateBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Println("CreateBlog func was invoked!", in)
	// implement struct BlogItem
	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	// insert data to mongo
	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Error(
			codes.Internal,
			fmt.Sprintf("mongo internal error: %v \n", err),
		)
	}

	oid, ok := res.InsertedID.(primitive.ObjectID)

	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			"cannot convert to OID",
		)
	}

	return &pb.BlogId{Id: oid.Hex()}, nil
}
