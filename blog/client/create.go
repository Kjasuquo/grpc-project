package main

import (
	"context"
	pb "github.com/kjasuquo/grpc/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("createBlog was invoked")

	blog := &pb.Blog{
		AuthorId: "Clement",
		Title:    "My first blog",
		Content:  "content of the first blog",
	}

	res, err := c.CreatBlog(context.Background(), blog)
	if err != nil {
		log.Printf("Unexpected error: %v\n", err)
	}
	log.Printf("Blog has been created with id: %s\n", res.Id)

	return res.Id
}
