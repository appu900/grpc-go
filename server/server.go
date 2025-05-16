package main

import (
	"context"
	"errors"
	"fmt"
	proto "grpc/protoc"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	proto.UnimplementedExampleServer
}

func (s *Server) ServerReplay(c context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
	fmt.Println("recive request from client", req.Something)
	fmt.Println("hello from server")
	return &proto.HelloResponse{}, errors.New("")
}

func main() {
	listner, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}
	grpcServer := grpc.NewServer()
	proto.RegisterExampleServer(grpcServer, &Server{})
	reflection.Register(grpcServer)

	if e := grpcServer.Serve(listner); e != nil {
		panic(e)
	}
}
