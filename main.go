package main

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	protoapi "protoapi/proto"
)

func main() {

	port := ":8080"

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
	return

}
