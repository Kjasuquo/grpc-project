package main

import (
	pb "github.com/kjasuquo/grpc/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type Server struct {
	pb.PrimeServiceServer
}

var addr = "localhost:50051"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error listening to connection: %v\n", err)
	}

	log.Printf("listening on: %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterPrimeServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}