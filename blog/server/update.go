package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func (s *Server) UpdateBlog(ctx context.Context, in *pb.Blog) (*emptypb.Empty, error) {
	log.Printf("UpdateBlog invoke with: %v\n", in)

	oid, err := primitive.ObjectIDFromHex(in.Id)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "ID not valid")
	}

	data := &pb.Blog{
		AuthorId: in.AuthorId,
		Title:    in.Title,
		Content:  in.Content,
	}

	filter := bson.M{"_id": oid}

	res, err := collection.UpdateOne(ctx, filter, bson.M{"$set": data})
	if err != nil {
		return nil, status.Errorf(codes.Internal, "Could not update")
	}

	if res.MatchedCount == 0 {
		return nil, status.Errorf(codes.NotFound, "Cannot find blog with id")
	}

	return &emptypb.Empty{}, nil

}
