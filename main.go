package main

import (
	"context"
	"grpc-course-protobuf/pb/user"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type userService struct {
	user.UnimplementedUserServiceServer
}

func (us *userService) CreateUser(ctx context.Context, userRequest *user.User) (*user.CreateResponse, error) {
	log.Println("Received request to create user")
	return &user.CreateResponse{
		Message: "User created successfully",
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("Failed to listen on port 8081: %v", err)
	}

	serv := grpc.NewServer()

	user.RegisterUserServiceServer(serv, &userService{})

	reflection.Register(serv) // local reflection for dev and testing purposes
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
