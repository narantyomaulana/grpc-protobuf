package main

import (
	"context"
	"errors"
	"grpc-course-protobuf/pb/chat"
	"grpc-course-protobuf/pb/user"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
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

type chatService struct {
	chat.UnimplementedChatServiceServer
}

func (cs *chatService) SendMessage(stream grpc.ClientStreamingServer[chat.ChatMessage, chat.ChatResponse]) error {

	for {
		req, err := stream.Recv()
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return status.Errorf(codes.Unknown, "Error Receiving message: %v", err)
		}

		log.Printf("Received Message from User ID %d: %s\n", req.UserId, req.Content)

	}

	return stream.SendAndClose(&chat.ChatResponse{
		Message: "Thank you for your message!",
	})

}

// func (UnimplementedChatServiceServer) ReceiveMessage(*ReceiveMessageRequest, grpc.ServerStreamingServer[ChatMessage]) error {
// 	return status.Errorf(codes.Unimplemented, "method ReceiveMessage not implemented")
// }
// func (UnimplementedChatServiceServer) Chat(grpc.BidiStreamingServer[ChatMessage, ChatMessage]) error {
// 	return status.Errorf(codes.Unimplemented, "method Chat not implemented")
// }

func main() {
	lis, err := net.Listen("tcp", ":8081")

	if err != nil {
		log.Fatalf("Failed to listen on port 8081: %v", err)
	}

	serv := grpc.NewServer()

	user.RegisterUserServiceServer(serv, &userService{})
	chat.RegisterChatServiceServer(serv, &chatService{})

	reflection.Register(serv) // local reflection for dev and testing purposes
	if err := serv.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
