package main

//
//import (
//	"context"
//	"fmt"
//	"log"
//	"net"
//)
//
//type server struct{}
//
//func (s *server) SayHello(ctx context.Context, req *server.HelloRequest) (*server.HelloResponse, error) {
//	return &server.HelloResponse{ResponseMsg: "Hello " + req.RequestName}, nil
//}
//
//func main() {
//	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 50051))
//	if err != nil {
//		log.Fatalf("failed to listen: %v", err)
//	}
//
//	s := grpc.NewServer()
//	server.RegisterSayHelloServer(s, &server{})
//	if err := s.Serve(lis); err != nil {
//		log.Fatalf("failed to serve: %v", err)
//	}
//}
