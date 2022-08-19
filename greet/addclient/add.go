package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/greet/proto"
	"log"
)

func doAdd(d pb.AddServiceClient) {
	log.Println("doAdd was invoked")

	ans, err := d.Add(context.Background(), &pb.AddRequest{
		FirstNum:  3,
		SecondNum: 10,
	})
	if err != nil {
		log.Fatalf("could not add: %v\n", err)
	}

	log.Printf("ANSWER: %v\n", ans.Answer)
}
