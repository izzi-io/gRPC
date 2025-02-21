package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/izzi-io/gRPC/greet/proto"
)

var addr string = "0.0.0.0:5051"

type Server struct {
	pb.GreetServiceServer
}

func main() {

	listen, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("shit hit the fan. can't listen")
	}
	defer listen.Close()

	log.Println("listening on", addr)

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})

	if err = s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
