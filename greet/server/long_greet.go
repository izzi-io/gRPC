package main

import (
	pb "github.com/izzi-io/gRPC/greet/proto"
	"google.golang.org/grpc"
	"io"
	"log"
)

func (s *Server) LongGreet(stream grpc.ClientStreamingServer[pb.GreetRequest, pb.GreetResponse]) error {
	log.Println("Got a client stream")

	res := ""
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&pb.GreetResponse{
				Result: res,
			})
		}
		if err != nil {
			log.Fatalf("Failed to receive a message: %v", err)
		}
		log.Printf("received request: %v\n", req)
		res += "Hello " + req.FirstName + "!\n"
	}
}
