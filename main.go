package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
	client2 "protoapi/client"
	protoapi "protoapi/proto"
)

func main() {

	port := ":8080"

	_type := flag.String("type", "client", "Server or Client ?")
	flag.Parse()

	if *_type == "client" {
		conn, err := grpc.NewClient("localhost" + port)
		if err != nil {
			log.Fatalf("could not connect: %v", err)
			return
		}

		client := protoapi.NewRandomClient(conn)
		r, err := client2.GetDateTime(context.Background(), client)

		if err != nil {
			log.Fatalf("cannot get datetime: %v", err)
			return
		}

		fmt.Println("Server datetime: ", r.Value)

		length := int64(rand.Intn(20))
		p, err := client2.GetPassword(context.Background(), client, 100, length+1)
		if err != nil {
			log.Fatalf("cannot get password: %v", err)
		}

		fmt.Println("Server password: ", p)

		place := int64(rand.Intn(100))
		seed := int64(rand.Intn(100))

		i, err := client2.GetRandom(context.Background(), client, seed, place)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Server random number: ", i.Value)

	} else {
		server := grpc.NewServer()
		var randomServer protoapi.RandomServer
		protoapi.RegisterRandomServer(server, randomServer)
		reflection.Register(server)

		listen, err := net.Listen("tcp", port)
		if err != nil {
			log.Fatalf("failed to listen: %v", err)
			return
		}

		fmt.Printf("listening on port %s\n", port)
		server.Serve(listen)
	}

	return
}
