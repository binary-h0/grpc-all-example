package main

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/binary-h0/grpc-all-example/ex08/proto"
	"google.golang.org/grpc"
)

type coordinatorServer struct {
	pb.UnimplementedCoordinatorServiceServer
	participants []pb.ParticipantServiceClient
}

func (s *coordinatorServer) StartTransaction(ctx context.Context, req *pb.TransactionRequest) (*pb.TransactionResponse, error) {
	var wg sync.WaitGroup
	statusChan := make(chan string, len(s.participants))
	rollbackChan := make(chan bool, 1)

	// Phase 1: Prepare
	for _, participant := range s.participants {
		wg.Add(1)
		go func(client pb.ParticipantServiceClient) {
			defer wg.Done()
			res, err := client.Prepare(ctx, req)
			if err != nil || res.Status != "PREPARED" {
				rollbackChan <- true
				return
			}
			statusChan <- "PREPARED"
		}(participant)
	}

	wg.Wait()
	close(statusChan)

	// Check if rollback is required
	select {
	case <-rollbackChan:
		// Rollback all participants
		for _, participant := range s.participants {
			_, _ = participant.Rollback(ctx, req)
		}
		return &pb.TransactionResponse{
			TransactionId: req.TransactionId,
			Status:        "ROLLED_BACK",
		}, nil
	default:
	}

	// Phase 2: Commit
	for _, participant := range s.participants {
		_, err := participant.Commit(ctx, req)
		if err != nil {
			return &pb.TransactionResponse{
				TransactionId: req.TransactionId,
				Status:        "FAILED",
			}, nil
		}
	}

	return &pb.TransactionResponse{
		TransactionId: req.TransactionId,
		Status:        "COMMITTED",
	}, nil
}

func main() {
	conn1, _ := grpc.Dial("localhost:50058", grpc.WithInsecure())
	conn2, _ := grpc.Dial("localhost:50059", grpc.WithInsecure())

	participants := []pb.ParticipantServiceClient{
		pb.NewParticipantServiceClient(conn1),
		pb.NewParticipantServiceClient(conn2),
	}

	server := grpc.NewServer()
	pb.RegisterCoordinatorServiceServer(server, &coordinatorServer{participants: participants})

	lis, _ := net.Listen("tcp", ":50057")
	log.Println("Coordinator server running on port 50057...")
	server.Serve(lis)
}
