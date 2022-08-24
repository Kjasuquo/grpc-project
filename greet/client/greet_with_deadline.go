package main

import (
	"context"
	"fmt"
	pb "github.com/kjasuquo/grpc/greet/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log"
	"time"
)

func doGreetWithDeadline(c pb.GreetServiceClient, timeout time.Duration) {
	log.Println("doGreetWithDeadline was invoked")

	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	fmt.Printf("this is contextttt: %v\n", ctx)
	defer cancel()

	req := &pb.GreetRequest{FirstName: "Joseph"}

	res, err := c.GreetWithDeadline(ctx, req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				log.Println("Deadline was exceeded")
				return
			} else {
				log.Fatalf("Unexpected error: %v\n", err)
			}
		} else {
			log.Fatalf("A non gRPC error: %v\n", err)
		}
	}

	fmt.Printf("Greet with deadline: %v\n", res.Result)
}
