package main

import (
	"context"
	pb "github.com/izzi-io/gRPC/calculator/proto"
	"log"
	"time"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	log.Printf("Received: %v", time.Now())
	return &pb.SumResponse{Result: in.FirstNumber + in.SecondNumber}, nil
}
