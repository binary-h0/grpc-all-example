package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"strings"
	"sync"

	pb "github.com/binary-h0/grpc-all-example/ex07/proto"
	"github.com/go-redis/redis/v8"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

type chatServer struct {
	pb.UnimplementedChatServiceServer
	redisClient *redis.Client
	subscribers map[string]chan *pb.Message
	mu          sync.Mutex
}

func newChatServer() *chatServer {
	client := redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})
	return &chatServer{
		redisClient: client,
		subscribers: make(map[string]chan *pb.Message),
	}
}

func (s *chatServer) SendMessage(ctx context.Context, msg *pb.Message) (*pb.Ack, error) {
	err := s.redisClient.Publish(ctx, "chat_channel", msg.Username+" "+msg.Content).Err()
	if err != nil {
		return nil, err
	}
	return &pb.Ack{Message: "Message sent successfully"}, nil
}

func (s *chatServer) ReceiveMessages(_ *pb.Empty, stream pb.ChatService_ReceiveMessagesServer) error {
	ch := make(chan *pb.Message, 100)

	md, ok := metadata.FromIncomingContext(stream.Context())
	if !ok {
		return fmt.Errorf("metadata is not set")
	}

	clientIDValue := md["clientid"]
	log.Printf("ClientID: %v", clientIDValue)

	if len(clientIDValue) == 0 {
		return fmt.Errorf("clientID is not set")
	}

	clientID := clientIDValue[0]

	s.mu.Lock()
	s.subscribers[clientID] = ch
	s.mu.Unlock()

	defer func() {
		s.mu.Lock()
		delete(s.subscribers, clientID)
		s.mu.Unlock()
		close(ch)
	}()

	for {
		select {
		case <-stream.Context().Done():
			return nil
		case msg := <-ch:
			if err := stream.Send(msg); err != nil {
				return err
			}
		}
	}
}

func (s *chatServer) subscribeRedis(ctx context.Context) {
	pubsub := s.redisClient.Subscribe(ctx, "chat_channel")
	defer pubsub.Close()

	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			log.Printf("Error subscribing to Redis: %v", err)
			return
		}

		data := strings.Split(msg.Payload, " ")
		username := data[0]
		text := strings.Join(data[1:], " ")

		message := &pb.Message{
			Username: username,
			Content:  text,
		}

		s.mu.Lock()
		for _, subscriber := range s.subscribers {
			select {
			case subscriber <- message:
			default:
			}
		}
		s.mu.Unlock()
	}
}

func main() {
	server := newChatServer()
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	go server.subscribeRedis(ctx)

	lis, err := net.Listen("tcp", ":50057")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterChatServiceServer(grpcServer, server)

	go func() {
		log.Println("Server is running on port 50057...")
		if err := grpcServer.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// Graceful shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	log.Println("Shutting down server...")
	grpcServer.GracefulStop()
}
