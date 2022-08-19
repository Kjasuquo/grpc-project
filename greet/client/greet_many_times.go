package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	stream, err := c.GreetManyTime(context.Background(), &pb.GreetRequest{FirstName: "Victor"})
	if err != nil {
		log.Fatalf("could not greet many times: %v\n", err)
	}

	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("err reading the stream %v\n", err)
		}
		log.Printf("Greet many times: %v\n", msg.Result)
	}
}
