package main

import (
	"log"
	"net"

	pb "github.com/grpcserver/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:50051"

type Server struct {
	pb.BookServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed %v", err)
	}

	log.Printf("Listening at: %s \n", addr)

	s := grpc.NewServer()
	pb.RegisterBookServiceServer(s, &Server{})

	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
