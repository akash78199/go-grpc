package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "grpc_user_service/proto/user"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Example: Fetch a user by ID
	user, err := client.GetUserById(context.Background(), &pb.UserRequest{Id: 1})
	if err != nil {
		log.Fatalf("Failed to fetch user: %v", err)
	}
	fmt.Printf("User by ID: %+v\n", user)

	// Example: Fetch a list of users by IDs
	userIds := []int32{1, 2}
	stream, err := client.GetUsersByIds(context.Background(), &pb.UserIdsRequest{Ids: userIds})
	if err != nil {
		log.Fatalf("Failed to fetch users by IDs: %v", err)
	}

	for {
		user, err := stream.Recv()
		if err != nil {
			break
		}
		fmt.Printf("User by IDs: %+v\n", user)
	}

	time.Sleep(time.Second)
}
