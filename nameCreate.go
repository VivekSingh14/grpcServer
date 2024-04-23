package main

import (
	"context"
	"log"

	pb "github.com/grpcserver/proto"
)

func (s *Server) NameCreated(ctx context.Context, in *pb.AuthRequest) (*pb.AuthResponse, error) {
	log.Printf(" NameCreated fucntion was invoked with %v \n ", in)
	return &pb.AuthResponse{
		BookName: in.Name + "-pvt",
	}, nil
}
