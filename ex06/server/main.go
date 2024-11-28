package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/binary-h0/grpc-all-example/ex06/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type productServer struct {
	pb.UnimplementedProductServiceServer
}

func (s *productServer) GetProduct(ctx context.Context, req *pb.GetProductRequest) (*pb.GetProductResponse, error) {
	return &pb.GetProductResponse{
		Id:    req.Id,
		Name:  "Sample Product",
		Price: 99.99,
	}, nil
}

func (s *productServer) ListProducts(req *pb.ListProductsRequest, stream pb.ProductService_ListProductsServer) error {
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second) // Simulate delay
		stream.Send(&pb.ListProductsResponse{
			Id:    fmt.Sprintf("prod-%d", i),
			Name:  fmt.Sprintf("Product %d", i),
			Price: float64(i) * 10,
		})
	}
	return nil
}

// Unary Interceptor for logging
func loggingUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	start := time.Now()
	log.Printf("Unary RPC: %s, Request: %v", info.FullMethod, req)

	resp, err := handler(ctx, req)
	if err != nil {
		log.Printf("Unary RPC Error: %v", err)
	} else {
		log.Printf("Unary RPC Response: %v, Duration: %s", resp, time.Since(start))
	}

	return resp, err
}

// Stream Interceptor for logging
func loggingStreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	start := time.Now()
	log.Printf("Streaming RPC: %s", info.FullMethod)

	err := handler(srv, ss)
	if err != nil {
		log.Printf("Streaming RPC Error: %v", err)
	} else {
		log.Printf("Streaming RPC Completed, Duration: %s", time.Since(start))
	}

	return err
}

// Unary Interceptor for authentication
func authUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok || len(md["authorization"]) == 0 || md["authorization"][0] != "Bearer valid-token" {
		return nil, fmt.Errorf("unauthorized")
	}

	return handler(ctx, req)
}

func main() {
	lis, err := net.Listen("tcp", ":50056")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	server := grpc.NewServer(
		grpc.ChainUnaryInterceptor(authUnaryInterceptor, loggingUnaryInterceptor),
		grpc.StreamInterceptor(loggingStreamInterceptor),
	)
	pb.RegisterProductServiceServer(server, &productServer{})

	log.Println("Server is running on port 50056...")

	if err := server.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
