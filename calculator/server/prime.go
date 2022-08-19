package main

import (
	pb "github.com/kjasuquo/grpc/calculator/proto"
	"log"
)

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.PrimeService_PrimeServer) error {
	log.Printf("Prime was invoked")

	var N = in.Number
	var k int32 = 2
	for N > 1 {
		if N%k == 0 { // if k evenly divides into N
			res := k // this is a factor

			err := stream.Send(&pb.PrimeResponse{Answer: res})
			if err != nil {
				log.Fatalf("error in streaming prim numbers: %v\n", err)
			}
			N = N / k // divide N by k so that we have the rest of the number left.
		} else {
			k = k + 1
		}

	}
	return nil
}
