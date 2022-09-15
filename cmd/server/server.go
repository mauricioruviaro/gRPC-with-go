package main

import (
	"log"
	"net"

	"github.com/mauricioruviaro/gRPC-with-go/services"
	"github.com/mauricioruviaro/gRPC-with-go/pb"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalf("could not connect: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, services.NewUserService())

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("could not serve: %v", err)
	}
}