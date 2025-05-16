package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	proto "grpc/protoc"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var client proto.ExampleClient

func main() {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect %v", err)
	}
	client = proto.NewExampleClient(conn)
	r := gin.Default()
	r.POST("/send-message", clientConnectionServer)
	r.Run(":8080")
}

func clientConnectionServer(c *gin.Context) {
	req := []*proto.HelloRequest{
		{Something: "Request1"},
		{Something: "Request2"},
		{Something: "Reqeust3"},
		{Something: "Request4"},
	}

	stream, err := client.ServerReplay(context.TODO())
	if err != nil {
		fmt.Println("Something went wrong !!")
		return
	}
	for _, re := range req {
		err := stream.Send(re)
		if err != nil {
			fmt.Println("request not fulfil")
			return
		}
	}

	response, err := stream.CloseAndRecv()
	if err != nil {
		fmt.Println("there is some error occure")
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message_count": response,
	})
}
