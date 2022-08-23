package main

import (
	"context"
	"fmt"
	pb "github.com/kjasuquo/grpc/greet/proto"
	"io"
	"log"
	"time"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")

	stream, err := c.GreetEveryone(context.Background())
	if err != nil {
		log.Fatalf("error while creating stream: %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Joseph"},
		{FirstName: "Augustine"},
		{FirstName: "Asuquo"},
	}

	waitChannel := make(chan struct{})

	go func() {
		for _, req := range reqs {
			err = stream.Send(req)
			if err != nil {
				log.Fatalf("error while sending stream: %v\n", err)
			}
			time.Sleep(1 * time.Second)
		}

		err = stream.CloseSend()
		if err != nil {
			log.Fatalf("error while closing sent stream: %v\n", err)
		}
	}()

	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}

			if err != nil {
				fmt.Printf("Error while receiving: %v\n", err)
				break
			}
			log.Printf("Recieved: %v\n", res.Result)
		}

		close(waitChannel)
	}()

	<-waitChannel
}
