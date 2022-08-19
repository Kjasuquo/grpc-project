package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/calculator/proto"
	"io"
	"log"
)

func doPrime(c pb.PrimeServiceClient) {
	log.Println("doPrime was invoked")

	stream, err := c.Prime(context.Background(), &pb.PrimeRequest{Number: 210})
	if err != nil {
		log.Fatalf("could not get prime number: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("err reading the stream %v\n", err)
		}

		log.Printf("prime number: %v\n", msg.Answer)
	}

}
