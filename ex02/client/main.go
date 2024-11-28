package main

import (
	"bufio"
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/binary-h0/grpc-all-example/ex02/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	// Chat 서비스 호출
	stream, err := client.Chat(context.Background())
	if err != nil {
		log.Fatalf("Failed to call Chat: %v", err)
	}

	// 수신 처리
	go func() {
		for {
			in, err := stream.Recv()
			if err != nil {
				log.Printf("Error receiving message: %v", err)
				return
			}
			fmt.Printf("[%s]: %s\n", in.User, in.Message)
		}
	}()

	// 입력 처리
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Enter your username:")
	scanner.Scan()
	username := scanner.Text()

	fmt.Println("Start chatting! Type your message and press Enter.")
	for scanner.Scan() {
		msg := scanner.Text()
		if err := stream.Send(&pb.ChatMessage{User: username, Message: msg}); err != nil {
			log.Printf("Error sending message: %v", err)
			return
		}
		time.Sleep(10 * time.Millisecond) // 속도 제한
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading input: %v", err)
	}
}
