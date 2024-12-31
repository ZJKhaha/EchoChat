package main

import (
	"context"
	"fmt"
	pb "github.com/ZJKhaha/EchoChat/hello-server/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	//pb.UnsafeSayHelloServer
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloResquest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{ResponseMsg: "Hello " + req.RequestName}, nil
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
