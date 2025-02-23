package main

import (
	"context"
	pb "github.com/izzi-io/gRPC/greet/proto"
	"log"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("do greet called")
	req := &pb.GreetRequest{
		FirstName: "izzi-io",
	}

	res, err := c.Greet(context.Background(), req)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	log.Println(res.Result)

}
