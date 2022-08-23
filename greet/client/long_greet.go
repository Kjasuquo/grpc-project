package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("doGreetManyTimes was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Tony"},
		{FirstName: "Anthony"},
		{FirstName: "Xavier"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("error while calling longGreet: %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("sending a req: %v\n", req)
		err = stream.Send(req)
		if err != nil {
			log.Fatalf("error in sending steams in client: %v\n", err)
		}
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("error while receiving response from longGreet: %v\n", err)
	}

	log.Printf("Long Greet: %s\n", res.Result)
}
