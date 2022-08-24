package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
)

var collection *mongo.Collection

type Server struct {
	pb.BlogServiceServer
}

var addr = "localhost:50051"

func main() {

	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://root:root@localhost:27017/"))
	if err != nil {
		log.Fatalf("error getting mongoDB client: %v\n", err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		log.Fatalf("error connecting to mongoDB client: %v\n", err)
	}

	collection = client.Database("blogdb").Collection("blog")

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("error listening to connection: %v\n", err)
	}

	log.Printf("listening on: %s\n", addr)

	s := grpc.NewServer()

	pb.RegisterBlogServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}

}
