package main

import (
	"context"
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
	r.POST("/sent-message/:message", clientConnectionServer)
	r.Run(":8080")
}

func clientConnectionServer(c *gin.Context) {
	message := c.Param("message")
	req := &proto.HelloRequest{Something: message}
	client.ServerReplay(context.TODO(), req)
	c.JSON(http.StatusOK, gin.H{
		"message": "message sent sucessfully to server",
	})
}
