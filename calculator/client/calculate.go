package main

import (
	"context"
	pb "github.com/izzi-io/gRPC/calculator/proto"
	"log"
)

func getSum(s pb.CalculatorServiceClient) {

	log.Println("get sum called")
	req := &pb.SumRequest{
		FirstNumber:  10,
		SecondNumber: 20,
	}

	response, err := s.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Sum error: %v", err)
	}
	log.Println(response.Result)

}
