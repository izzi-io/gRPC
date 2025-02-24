package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	"log"

	pb "github.com/izzi-io/gRPC/greet/proto"
)

var addr string = "localhost:5051"

func main() {
	conn, err := grpc.NewClient(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect: %v", err)
		return
	}
	defer conn.Close()

	c := pb.NewGreetServiceClient(conn)

	greetManyTimes(c)

}
