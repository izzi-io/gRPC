package main

import (
	"context"
	pb "github.com/izzi-io/gRPC/greet/proto"
	"io"
	"log"
)

func greetManyTimes(client pb.GreetServiceClient) {

	log.Println("start greetManyTimes")
	request := &pb.GreetRequest{
		FirstName: "izzi-io",
	}

	stream, err := client.GreetManyTimes(context.Background(), request)

	if err != nil {
		log.Fatalf("%v.GreetManyTimes RPC call failed.", err)
		return
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Fatalf("%v.GreetManyTimes RPC call failed while reading the stream", err)
		}

		log.Println(msg)

	}

}
