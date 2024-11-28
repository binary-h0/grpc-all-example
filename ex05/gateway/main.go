package main

import (
	"context"
	"log"
	"net/http"

	pb "github.com/binary-h0/grpc-all-example/ex05/proto"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
)

func main() {
	grpcEndpoint := "localhost:50055"
	httpEndpoint := ":8080"

	conn, err := grpc.Dial(grpcEndpoint, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to dial gRPC server: %v", err)
	}
	defer conn.Close()

	mux := runtime.NewServeMux()
	err = pb.RegisterUserServiceHandler(context.Background(), mux, conn)
	if err != nil {
		log.Fatalf("Failed to register gRPC Gateway: %v", err)
	}

	log.Printf("REST Gateway running on %s...", httpEndpoint)
	log.Fatal(http.ListenAndServe(httpEndpoint, mux))
}
