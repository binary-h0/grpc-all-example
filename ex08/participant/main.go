package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/binary-h0/grpc-all-example/ex08/proto"
	"google.golang.org/grpc"
)

type participantServer struct {
	pb.UnimplementedParticipantServiceServer
	storage map[string]string
}

func (s *participantServer) Prepare(ctx context.Context, req *pb.TransactionRequest) (*pb.StatusResponse, error) {
	log.Printf("Preparing transaction: %s", req.TransactionId)
	s.storage[req.TransactionId] = "PREPARED"
	return &pb.StatusResponse{Status: "PREPARED"}, nil
}

func (s *participantServer) Commit(ctx context.Context, req *pb.TransactionRequest) (*pb.StatusResponse, error) {
	log.Printf("Committing transaction: %s", req.TransactionId)
	s.storage[req.TransactionId] = "COMMITTED"
	return &pb.StatusResponse{Status: "COMMITTED"}, nil
}

func (s *participantServer) Rollback(ctx context.Context, req *pb.TransactionRequest) (*pb.StatusResponse, error) {
	log.Printf("Rolling back transaction: %s", req.TransactionId)
	delete(s.storage, req.TransactionId)
	return &pb.StatusResponse{Status: "ROLLED_BACK"}, nil
}

func main() {
	// 플래그 설정
	port := flag.String("port", "50058", "Port to run the server on")
	flag.Parse() // 플래그 파싱

	server := grpc.NewServer()
	pb.RegisterParticipantServiceServer(server, &participantServer{storage: make(map[string]string)})

	address := fmt.Sprintf(":%s", *port)
	lis, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("Failed to listen on port %s: %v", *port, err)
	}

	log.Printf("Participant server running on port %s...", *port)
	server.Serve(lis)
}
