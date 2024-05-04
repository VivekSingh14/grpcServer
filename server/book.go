package main

import (
	"context"
	"log"

	pb "github.com/grpcserver/proto"
)

func (s *Server) Book(ctx context.Context, in *pb.BookRequest) (*pb.BookResponse, error) {
	log.Printf(" BookCreated fucntion was invoked with %v \n ", in)
	return &pb.BookResponse{
		BookName: in.Name + "-pvt",
	}, nil
}
