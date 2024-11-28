package main

import (
	"context"
	"log"

	pb "github.com/binary-h0/grpc-all-example/ex06/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func main() {
	conn, err := grpc.Dial("localhost:50056", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewProductServiceClient(conn)

	ctx := metadata.AppendToOutgoingContext(context.Background(), "authorization", "Bearer valid-token")

	// Call GetProduct
	res, err := client.GetProduct(ctx, &pb.GetProductRequest{Id: "prod-1"})
	if err != nil {
		log.Fatalf("GetProduct failed: %v", err)
	}
	log.Printf("GetProduct Response: %v", res)

	// Call ListProducts
	stream, err := client.ListProducts(ctx, &pb.ListProductsRequest{})
	if err != nil {
		log.Fatalf("ListProducts failed: %v", err)
	}

	for {
		item, err := stream.Recv()
		if err != nil {
			break
		}
		log.Printf("Product: %v", item)
	}
}
