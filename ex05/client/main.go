package main

import (
	"context"
	"log"

	pb "github.com/binary-h0/grpc-all-example/ex05/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50055", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewUserServiceClient(conn)

	// Create a user
	createRes, err := client.CreateUser(context.Background(), &pb.CreateUserRequest{
		Name: "John Doe",
		Age:  30,
	})
	if err != nil {
		log.Fatalf("CreateUser failed: %v", err)
	}
	log.Printf("Created User ID: %s", createRes.Id)

	// Get user information
	getRes, err := client.GetUser(context.Background(), &pb.GetUserRequest{Id: createRes.Id})
	if err != nil {
		log.Fatalf("GetUser failed: %v", err)
	}
	log.Printf("User Info: ID=%s, Name=%s, Age=%d", getRes.Id, getRes.Name, getRes.Age)
}
