package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/binary-h0/grpc-all-example/ex03/proto"
	"github.com/golang-jwt/jwt/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type authServer struct {
	pb.UnimplementedAuthServiceServer
}

var jwtKey = []byte("my_secret_key")

func (s *authServer) Login(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	if req.Username == "user" && req.Password == "password" {
		// Generate JWT token
		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"username": req.Username,
		})
		tokenString, err := token.SignedString(jwtKey)
		if err != nil {
			return nil, err
		}
		return &pb.LoginResponse{Token: tokenString}, nil
	}
	return nil, fmt.Errorf("invalid credentials")
}

func (s *authServer) AccessProtectedData(ctx context.Context, req *pb.AccessRequest) (*pb.AccessResponse, error) {
	token, err := jwt.Parse(req.Token, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	return &pb.AccessResponse{Data: "This is protected data!"}, nil
}

func main() {
	// Load TLS credentials
	creds, err := credentials.NewServerTLSFromFile("certs/server.crt", "certs/server.key")
	if err != nil {
		log.Fatalf("Failed to load TLS credentials: %v", err)
	}

	lis, err := net.Listen("tcp", ":50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.Creds(creds))
	pb.RegisterAuthServiceServer(s, &authServer{})

	log.Println("Auth server running on port 50053...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
