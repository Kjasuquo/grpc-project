package main

import (
	"context"
	"fmt"
	pb "github.com/kjasuquo/grpc/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"math"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Println("SqRt function invoked")

	number := in.Number

	if number < 0 {
		return nil, status.Errorf(
			codes.InvalidArgument,
			fmt.Sprintf("Recieved a negative number: %d\n", number),
		)
	}

	return &pb.SqrtResponse{Result: math.Sqrt(float64(number))}, nil
}
