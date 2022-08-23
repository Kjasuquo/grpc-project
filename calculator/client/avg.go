package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/calculator/proto"
	"log"
	"time"
)

func doAvg(c pb.PrimeServiceClient) {
	log.Println("doAvg function in the client side invoked")

	number := []int32{3, 9, 5, 32}
	stream, err := c.Avg(context.Background())
	if err != nil {
		log.Fatalf("error while openning the stream: %v\n", err)
	}
	for _, req := range number {
		log.Printf("sending req: %v\n", req)
		err = stream.Send(&pb.AvgRequest{Number: req})
		if err != nil {
			log.Fatalf("error in sending steams in client: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error in receiving steams in client: %v\n", err)
	}

	log.Printf("Average: %v\n", res.Result)
}
