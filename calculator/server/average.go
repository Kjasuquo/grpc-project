package main

import (
	pb "github.com/kjasuquo/grpc/calculator/proto"
	"io"
	"log"
)

func (s *Server) Average(stream pb.PrimeService_AverageServer) error {
	log.Println("Average function was invoked")

	var sum int32
	var n int32
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.PrimeResponse{Answer: sum / n})
		}
		if err != nil {
			log.Fatalf("error while reading Average client stream: %v\n", err)
		}

		sum += req.Number
		n = n + 1

	}
}
