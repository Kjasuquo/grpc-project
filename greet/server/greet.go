package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/greet/proto"
	"log"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Greet Func was invoked with: %v\n", in)

	return &pb.GreetResponse{Result: "Hello " + in.FirstName}, nil

}
