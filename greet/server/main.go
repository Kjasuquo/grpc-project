package main

import (
	pb "github.com/kjasuquo/grpc/greet/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr = "0.0.0.0:50051"

// Server is a grpc server object
type Server struct {
	pb.GreetServiceServer
	pb.AddServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen on: %v\n", err)
	}

	log.Printf("listening on: %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})
	pb.RegisterAddServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
