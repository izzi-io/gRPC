package main

import (
	"context"
	pb "github.com/izzi-io/gRPC/greet/proto"
	"log"
	"time"
)

func (s *Server) Greet(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {
	log.Printf("Received: %v", time.Now())
	return &pb.GreetResponse{Result: "Hello " + in.FirstName + " at " + time.Now().String()}, nil
}
