package main

import (
	"log"
	"math/rand"
	"time"

	"net"

	pb "github.com/binary-h0/grpc-all-example/ex09/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedStreamingServiceServer
}

func (s *server) SendData(req *pb.Empty, stream pb.StreamingService_SendDataServer) error {
	for {
		data := &pb.SensorData{
			SensorId:    "sensor-1",
			Temperature: 20 + rand.Float64()*10, // Random temperature
			Humidity:    30 + rand.Float64()*20, // Random humidity
			Timestamp:   time.Now().Unix(),
		}

		if err := stream.Send(data); err != nil {
			return err
		}

		log.Printf("Sent data: %+v", data)
		time.Sleep(1 * time.Second) // Simulate data interval
	}
}

func main() {
	lis, err := net.Listen("tcp", ":50060")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterStreamingServiceServer(s, &server{})

	log.Println("Streaming server is running on port 50060...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
