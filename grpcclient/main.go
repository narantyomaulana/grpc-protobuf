package main

import (
	"context"
	"grpc-course-protobuf/pb/user"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	clientConn, err := grpc.NewClient("localhost:8081", grpc.WithTransportCredentials(insecure.NewCredentials())) // Create a gRPC client connection to the server running on localhost:8081
	if err != nil {
		log.Fatal("Failed to create gRPC client connection:", err)
	}

	userClient := user.NewUserServiceClient(clientConn)

	response, err := userClient.CreateUser(context.Background(), &user.User{
		Id:      1,
		Age:     13,
		Balance: 13000,
		Address: &user.Address{
			Id:          123,
			FullAddress: "123 Main St, Springfield",
			Province:    "Springfield Province",
			City:        "Springfield City",
		},
	})

	if err != nil {
		log.Fatal("Failed to create user:", err)
	} else {
		log.Println("Response from server:", response.Message)
	}
}
