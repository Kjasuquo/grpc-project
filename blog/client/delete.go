package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/blog/proto"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, id string) {
	log.Printf("deleteBlog was invoked with in: %v\n", id)

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: id})
	if err != nil {
		log.Fatalf("Error happened while deleting: %v\n", err)
	}

	log.Println("Blog was successfully deleted")

}
