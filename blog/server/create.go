package main

import (
	"context"
	"fmt"
	pb "github.com/kjasuquo/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
)

func (s *Server) CreatBlog(ctx context.Context, in *pb.Blog) (*pb.BlogId, error) {
	log.Println("CreateBlog invoked")

	data := BlogItem{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	res, err := collection.InsertOne(ctx, data)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Internal Error: %v\n", err))
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(codes.Internal, "Cannot convert to oid")
	}

	return &pb.BlogId{Id: oid.Hex()}, nil
}
