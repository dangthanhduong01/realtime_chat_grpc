package main

import (
	"fmt"
	"log"
	"net"
	pb "snowApp/gen"

	"google.golang.org/grpc"
)

type chatServer struct {
	pb.UnimplementedChatServiceServer
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &chatServer{})
	fmt.Println("Server is running on port :50051")
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
