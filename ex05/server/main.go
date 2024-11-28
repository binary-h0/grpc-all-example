package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/binary-h0/grpc-all-example/ex05/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type userServer struct {
	pb.UnimplementedUserServiceServer
	users map[string]*pb.GetUserResponse
}

func (s *userServer) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	id := fmt.Sprintf("user-%d", len(s.users)+1)
	s.users[id] = &pb.GetUserResponse{
		Id:   id,
		Name: req.Name,
		Age:  req.Age,
	}
	return &pb.CreateUserResponse{Id: id}, nil
}

func (s *userServer) GetUser(ctx context.Context, req *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	user, exists := s.users[req.Id]
	if !exists {
		return nil, fmt.Errorf("user not found")
	}
	return user, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, &userServer{users: make(map[string]*pb.GetUserResponse)})
	reflection.Register(grpcServer)

	log.Println("gRPC server is running on port 50055...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
