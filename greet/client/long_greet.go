package main

import (
	"context"
	pb "github.com/izzi-io/gRPC/greet/proto"
	"log"
	"time"
)

func longGreet(c pb.GreetServiceClient) {
	log.Printf("Call longGreet service")
	reqs := []*pb.GreetRequest{
		{FirstName: "izzi-io_1"},
		{FirstName: "izzi-io_2"},
		{FirstName: "izzi-io_3"},
		{FirstName: "izzi-io_4"},
		{FirstName: "izzi-io_5"},
		{FirstName: "izzi-io_6"},
		{FirstName: "izzi-io_7"},
	}

	stream, err := c.LongGreet(context.Background())
	if err != nil {
		log.Fatalf("%v.LongGreet(_) = _, %v", c, err)
	}

	for _, req := range reqs {
		if err := stream.Send(req); err != nil {
			log.Fatalf("%v.LongGreet(_) = _, %v", c, err)
		}
		time.Sleep(1000 * time.Millisecond)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.LongGreet(_) = _, %v", c, err)
	}

	log.Printf("Response from LongGreet: %s", res.Result)

}
