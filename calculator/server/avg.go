package main

import (
	pb "github.com/kjasuquo/grpc/calculator/proto"
	"io"
	"log"
)

func (s *Server) Avg(stream pb.PrimeService_AvgServer) error {
	log.Println("Avg is invoked")

	sum := 0
	count := 0

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.AvgResponse{Result: float64(sum) / float64(count)})
		}
		if err != nil {
			log.Fatalf("error while reading Average client stream: %v\n", err)
		}

		sum += int(req.Number)
		count++
	}

}
