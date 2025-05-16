package main

import (
	"fmt"
	proto "grpc/protoc"
	"io"
	"net"
	"strconv"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Server struct {
	proto.UnimplementedExampleServer
}

// ** Normal server replay call no stram only once response

// func (s *Server) ServerReplay(c context.Context, req *proto.HelloRequest) (*proto.HelloResponse, error) {
// 	fmt.Println("recive request from client", req.Something)
// 	fmt.Println("hello from server")
// 	return &proto.HelloResponse{}, errors.New("")
// }


// server replay will recive data in the form of streams
func (s *Server) ServerReplay(stream proto.Example_ServerReplayServer) error {
	totalMessageCount := 0
	for {
		request, error := stream.Recv()
		if error == io.EOF {
			return stream.SendAndClose(&proto.HelloResponse{
				Reply: strconv.Itoa(totalMessageCount),
			})
		}
		if error != nil {
			return error
		}
		totalMessageCount++
		fmt.Println(request)
	}
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
