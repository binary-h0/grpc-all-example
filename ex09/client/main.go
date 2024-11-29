package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/binary-h0/grpc-all-example/ex09/proto"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50060", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewStreamingServiceClient(conn)

	stream, err := client.SendData(context.Background(), &pb.Empty{})
	if err != nil {
		log.Fatalf("Failed to call SendData: %v", err)
	}

	var count int
	var totalTemp, totalHumidity float64

	fmt.Println("Real-time Sensor Data Dashboard")
	fmt.Println("--------------------------------")

	for {
		data, err := stream.Recv()
		if err != nil {
			log.Fatalf("Error receiving data: %v", err)
		}

		count++
		totalTemp += data.Temperature
		totalHumidity += data.Humidity

		avgTemp := totalTemp / float64(count)
		avgHumidity := totalHumidity / float64(count)

		fmt.Printf("Sensor ID: %s | Temp: %.2f°C | Humidity: %.2f%% | Avg Temp: %.2f°C | Avg Humidity: %.2f%%\n",
			data.SensorId, data.Temperature, data.Humidity, avgTemp, avgHumidity)

		time.Sleep(1 * time.Second)
	}
}
