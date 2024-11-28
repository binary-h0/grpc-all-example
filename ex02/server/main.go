package main

import (
	"io"
	"log"
	"net"
	"sync"

	pb "github.com/binary-h0/grpc-all-example/ex02/proto"
	"google.golang.org/grpc"
)

type chatServer struct {
	pb.UnimplementedChatServiceServer
	mu      sync.Mutex
	clients []chan *pb.ChatMessage
}

func (s *chatServer) Chat(stream pb.ChatService_ChatServer) error {
	// 클라이언트용 메시지 채널 생성
	msgChan := make(chan *pb.ChatMessage, 10)
	defer close(msgChan)

	s.mu.Lock()
	s.clients = append(s.clients, msgChan)
	s.mu.Unlock()

	// 수신 및 송신 핸들링
	go func() {
		for msg := range msgChan {
			if err := stream.Send(msg); err != nil {
				log.Printf("Error sending message: %v", err)
				return
			}
		}
	}()

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			log.Println("Client disconnected")
			return nil
		}
		if err != nil {
			log.Printf("Error receiving message: %v", err)
			return err
		}

		// 모든 클라이언트에게 메시지 브로드캐스트
		s.mu.Lock()
		for _, client := range s.clients {
			client <- in
		}
		s.mu.Unlock()

		log.Printf("[%s]: %s", in.User, in.Message)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterChatServiceServer(s, &chatServer{})

	log.Println("Chat server is running on port 50052...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
