package main

import (
	"fmt"
	pb "github.com/izzi-io/gRPC/greet/proto"
	"google.golang.org/grpc"
	"log"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream grpc.ServerStreamingServer[pb.GreetResponse]) error {

	log.Printf("greet many times function called, ", in)
	for i := 0; i < 10; i++ {
		res := fmt.Sprintf("Hello, %d %s!", i, in.FirstName)

		stream.Send(&pb.GreetResponse{Result: res})
	}
	return nil
}
