package main

import (
	"fmt"
	pb "github.com/kjasuquo/grpc/greet/proto"
	"log"
)

func (s *Server) GreetManyTime(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimeServer) error {
	log.Printf("GreetManyTimes Func was invoked with: %v\n", in)

	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello %s, number %d", in.FirstName, i+1)

		err := stream.Send(&pb.GreetResponse{Result: res})
		if err != nil {
			log.Fatalf("error in streaming greet: %v\n", err)

		}
	}

	return nil

}
