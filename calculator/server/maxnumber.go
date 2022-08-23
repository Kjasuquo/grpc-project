package main

import (
	pb "github.com/kjasuquo/grpc/calculator/proto"
	"io"
	"log"
)

func (s *Server) Max(stream pb.PrimeService_MaxServer) error {
	log.Println("Max server invoked")
	var num int32 = 0
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v\n", err)
		}

		if req.Number > num {
			num = req.Number

			err = stream.Send(&pb.MaxResponse{Result: num})
			if err != nil {
				log.Fatalf("Error while sending response stream: %v\n", err)
			}
		}
	}
}
