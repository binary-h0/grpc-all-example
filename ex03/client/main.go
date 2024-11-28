package main

import (
	"context"
	"crypto/tls"
	"log"
	"time"

	pb "github.com/binary-h0/grpc-all-example/ex03/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

func main() {
	creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	conn, err := grpc.NewClient("localhost:50053", grpc.WithTransportCredentials(creds))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewAuthServiceClient(conn)

	// Login and get token
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.Login(ctx, &pb.LoginRequest{Username: "user", Password: "password"})
	if err != nil {
		log.Fatalf("Login failed: %v", err)
	}
	log.Printf("Token: %s", res.Token)

	// Access protected data
	dataRes, err := client.AccessProtectedData(ctx, &pb.AccessRequest{Token: res.Token})
	if err != nil {
		log.Fatalf("Access to protected data failed: %v", err)
	}
	log.Printf("Protected Data: %s", dataRes.Data)
}
