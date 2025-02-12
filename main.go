package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"math/rand"
	"net"
	pbclient "protoapi/client"
	protoapi "protoapi/proto"
	gserver "protoapi/server"
)

func main() {

	port := ":8080"

	_type := flag.String("type", "pbclient", "Server or Client ?")
	flag.Parse()

	if *_type == "client" {
		conn, err := grpc.NewClient("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("could not connect: %v", err)
			return
		}

		defer conn.Close()

		clt := protoapi.NewRandomClient(conn)

		r, err := pbclient.GetDateTime(context.Background(), clt)

		if err != nil {
			log.Fatalf("cannot get datetime: %v", err)
			return
		}

		fmt.Println("Server datetime: ", r.Value)

		length := int64(rand.Intn(20))
		p, err := pbclient.GetPassword(context.Background(), clt, 100, length+1)
		if err != nil {
			log.Fatalf("cannot get password: %v", err)
		}

		fmt.Println("Server password: ", p)

		place := int64(rand.Intn(100))
		seed := int64(rand.Intn(100))

		i, err := pbclient.GetRandom(context.Background(), clt, seed, place)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("Server random number: ", i.Value)

	} else {
		server := grpc.NewServer()
		var randomServer gserver.RandomServer
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
