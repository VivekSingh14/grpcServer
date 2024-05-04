package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/grpcserver/proto"
)

var addr string = "0.0.0.0:50052"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()

	c := pb.NewBookServiceClient(conn)

	doGet(c)
}

func doGet(c pb.BookServiceClient) {
	log.Println("doGet is invoked...")

	res, err := c.Book(context.Background(), &pb.BookRequest{
		Name: "Vivek",
	})

	if err != nil {
		log.Fatalf("Could not greet %v\n", err)
	}

	log.Printf("Name Created: %s\n", res.BookName)
}
