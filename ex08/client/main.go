package main

import (
	"context"
	"log"
	"time"

	pb "github.com/binary-h0/grpc-all-example/ex08/proto"
	"google.golang.org/grpc"
)

func main() {
	// Coordinator 서버와 연결
	conn, err := grpc.Dial("localhost:50057", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to Coordinator: %v", err)
	}
	defer conn.Close()

	client := pb.NewCoordinatorServiceClient(conn)

	// 트랜잭션 요청 생성
	transactionID := "txn-12345"
	request := &pb.TransactionRequest{
		TransactionId: transactionID,
		Data:          "Example transaction payload",
	}

	// Coordinator의 StartTransaction 호출
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	response, err := client.StartTransaction(ctx, request)
	if err != nil {
		log.Fatalf("StartTransaction failed: %v", err)
	}

	log.Printf("Transaction ID: %s, Status: %s", response.TransactionId, response.Status)
}
