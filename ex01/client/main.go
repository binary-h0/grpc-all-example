package main

import (
	"context"
	"log"
	"time"

	pb "github.com/binary-h0/grpc-all-example/ex01/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewExampleServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	for i := 0; i < 100; i++ {
		res, err := client.SayHello(ctx, &pb.HelloRequest{Name: "Binary H0"})
		if err != nil {
			log.Fatalf("Failed to call SayHello: %v", err)
		}
		log.Printf("Response: %s", res.Message)
	}
}
