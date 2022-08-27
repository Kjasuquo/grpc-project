package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string) {
	log.Println("updateBlog was invoked")

	newBlog := &pb.Blog{
		Id:       id,
		AuthorId: "Not Tony",
		Title:    "A new Title",
		Content:  "Content of the first blog updated",
	}

	_, err := c.UpdateBlog(context.Background(), newBlog)
	if err != nil {
		log.Fatalf("Error happened while updating: %v\n", err)
	}

	log.Println("Blog was updated")
}
