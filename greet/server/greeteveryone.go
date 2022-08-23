package main

import (
	pb "github.com/kjasuquo/grpc/greet/proto"
	"io"
	"log"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Println("Invoking Greet Everyone")

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"

		err = stream.Send(&pb.GreetResponse{Result: res})
		if err != nil {
			log.Fatalf("error while sending stream: %v\n", err)
		}
	}
}
