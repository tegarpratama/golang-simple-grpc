package main

import (
	"context"
	"log"
	"net/http"

	messageProto "project-one/cmd/proto/message"

	"github.com/gin-gonic/gin"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	r := gin.Default()

	r.POST("/send", func(c *gin.Context) {
		var input struct {
			Message string `json:"message"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("Failed to connect to gRPC server: %v", err)
		}
		defer conn.Close()

		client := messageProto.NewMessageServiceClient(conn)
		resp, err := client.SendMessage(context.Background(), &messageProto.MessageRequest{
			Content: input.Message,
		})

		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"status":  resp.Status,
			"message": resp.Content,
		})
	})

	r.Run(":8080")
}
