package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/calculator/proto"
	"log"
	"time"
)

func doAverage(c pb.PrimeServiceClient) {
	log.Println("Average function in the client side invoked")
	reqs := []*pb.PrimeRequest{
		{Number: 2},
		{Number: 4},
		{Number: 3},
	}

	stream, err := c.Average(context.Background())
	if err != nil {
		log.Fatalf("error while calling average in client: %v\n", err)
	}
	for _, req := range reqs {
		log.Printf("sending req: %v\n", req)
		err = stream.Send(req)
		if err != nil {
			log.Fatalf("error in sending steams in client: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error in receiving steams in client: %v\n", err)
	}

	log.Printf("Average: %v\n", res.Answer)
}
