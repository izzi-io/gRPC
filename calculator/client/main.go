package main

import (
	pb "github.com/izzi-io/gRPC/calculator/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

func main() {
	conn, err := grpc.NewClient("localhost:5051", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	defer conn.Close()

	conx := pb.NewCalculatorServiceClient(conn)

	getSum(conx)

}
