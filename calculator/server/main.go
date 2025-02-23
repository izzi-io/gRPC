package main

import (
	pb "github.com/izzi-io/gRPC/calculator/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

var addr string = "localhost:5051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	defer listen.Close()

	log.Printf("server listening at %v", addr)

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, new(Server))

	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
