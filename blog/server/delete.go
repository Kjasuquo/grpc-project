package main

import (
	"context"
	"fmt"
	pb "github.com/kjasuquo/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) DeleteBlog(ctx context.Context, in *pb.BlogId) (*emptypb.Empty, error) {
	log.Printf("DeleteBlog invoked with %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "cannot parse ID")
	}

	filter := primitive.M{"_id": oid}

	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		return nil, status.Errorf(codes.Internal, fmt.Sprintf("Cannot delete obj in mongoDB: %v\n", err))
	}

	if res.DeletedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "blog was not found")
	}

	log.Printf("%v successfully deleted", res)

	return &emptypb.Empty{}, nil

}
