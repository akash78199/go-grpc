package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc_user_service/proto/user"

	"google.golang.org/grpc"
)

var users = []pb.User{
	{Id: 1, Fname: "Steve", City: "LA", Phone: 1234567890, Height: 5.8, Married: true},
	{Id: 2, Fname: "Alice", City: "NYC", Phone: 9876543210, Height: 5.6, Married: false},
	// Add more user details here
}

type userServiceServer struct{}

func (s *userServiceServer) GetUserById(ctx context.Context, request *pb.UserRequest) (*pb.User, error) {
	// Iterate through the users to find the one with the requested ID
	for _, user := range users {
		if user.Id == request.Id {
			return &user, nil
		}
	}
	// If the user is not found, return an error
	return nil, fmt.Errorf("User with ID %d not found", request.Id)
}

func (s *userServiceServer) GetUsersByIds(request *pb.UserIdsRequest, stream pb.UserService_GetUsersByIdsServer) error {
	// Iterate through the requested IDs
	for _, id := range request.Ids {
		found := false
		// Find the user with the current ID
		for _, user := range users {
			if user.Id == id {
				// Send the user to the client
				if err := stream.Send(&user); err != nil {
					return err
				}
				found = true
				break
			}
		}
		// If the user is not found, send an error to the client
		if !found {
			return fmt.Errorf("User with ID %d not found", id)
		}
	}
	return nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterUserServiceServer(server, &userServiceServer{})

	fmt.Println("Server is running on port :50051")
	if err := server.Serve(listen); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
