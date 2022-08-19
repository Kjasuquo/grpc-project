package main

import (
	pb "github.com/kjasuquo/grpc/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr = "localhost:50051"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect %v\n", err)
	}

	defer conn.Close()

	c := pb.NewPrimeServiceClient(conn)

	doPrime(c)
}
