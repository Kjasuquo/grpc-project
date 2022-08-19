package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/greet/proto"
	"log"
)

func (s *Server) Add(ctx context.Context, in *pb.AddRequest) (*pb.AddResponse, error) {
	log.Printf("Add was invoked with: %v\n", in)

	return &pb.AddResponse{Answer: in.FirstNum + in.SecondNum}, nil
}
