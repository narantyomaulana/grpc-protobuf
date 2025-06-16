package main

import (
	"context"
	"grpc-course-protobuf/pb/chat"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials())) // Create a gRPC client connection to the server running on localhost:8081
	if err != nil {
		log.Fatal("Failed to create gRPC client connection:", err)
	}

	chatClient := chat.NewChatServiceClient(clientConn)
	stream, err := chatClient.SendMessage(context.Background())

	if err != nil {
		log.Fatal("Failed to create stream:", err)
	}

	err = stream.Send(&chat.ChatMessage{
		UserId:  123,
		Content: "Hello, this is a test message!",
	})

	if err != nil {
		log.Fatal("Failed to send message:", err)
	}

	err = stream.Send(&chat.ChatMessage{
		UserId:  123,
		Content: "Hello, Again!",
	})

	if err != nil {
		log.Fatal("Failed to send message:", err)
	}

	time.Sleep(5 * time.Second)

	err = stream.Send(&chat.ChatMessage{
		UserId:  123,
		Content: "Hello, Brother!",
	})

	if err != nil {
		log.Fatal("Failed to send message:", err)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatal("Failed to close :", err)
	}

	log.Println("Connection is close. Message: ", res.Message)
}
