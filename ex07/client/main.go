package main

import (
	"bufio"
	"context"
	"log"
	"os"

	pb "github.com/binary-h0/grpc-all-example/ex07/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func receiveMessages(client pb.ChatServiceClient, ctx context.Context) {
	stream, err := client.ReceiveMessages(ctx, &pb.Empty{})

	if err != nil {
		log.Fatalf("Failed to receive messages: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving message: %v", err)
		}
		log.Printf("[%s]: %s", msg.Username, msg.Content)
	}
}

func main() {
	conn, err := grpc.Dial("localhost:50057", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewChatServiceClient(conn)

	scanner := bufio.NewScanner(os.Stdin)
	log.Println("Start chatting!")

	log.Print("Enter your username: ")
	scanner.Scan()
	username := scanner.Text()
	md := metadata.Pairs("clientID", username)
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	go receiveMessages(client, ctx)

	for scanner.Scan() {
		text := scanner.Text()
		_, err := client.SendMessage(context.Background(), &pb.Message{
			Username: username,
			Content:  text,
		})
		if err != nil {
			log.Fatalf("Failed to send message: %v", err)
		}
	}
}
