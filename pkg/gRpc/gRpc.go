package gRpc

import (
	"net"
	"log"
	"google.golang.org/grpc"
)


func SetUpServer(port string) *grpc.Server{

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()

	log.Println("Running on port:", port)
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	return server
}
