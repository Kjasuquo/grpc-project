package main

import (
	"context"
	"fmt"
	pb "github.com/kjasuquo/grpc/calculator/proto"
	"io"
	"log"
	"time"
)

func doMax(c pb.PrimeServiceClient) {
	log.Println("doMax invoked")

	stream, err := c.Max(context.Background())
	if err != nil {
		log.Fatalf("error while creating stream: %v\n", err)
	}

	number := []int32{1, 5, 3, 6, 2, 20}

	waitChannel := make(chan struct{})

	go func() {
		for _, req := range number {
			err = stream.Send(&pb.MaxRequest{Number: req})
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
