package main

import (
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/CDonosoK/T2-Distribuidos/chat"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Error listening on port 9000: %v", err)
	}

	s := chat.Server{}

	grpcServer := grpc.NewServer()

	chat.RegisterChatServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Error starting gRPC server on port 9000: %v", err)
	}
}
