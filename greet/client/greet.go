package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{FirstName: "Tony"})
	if err != nil {
		log.Fatalf("could not greet: %v\n", err)
	}

	log.Printf("GREETING %v\n", res.Result)
}
