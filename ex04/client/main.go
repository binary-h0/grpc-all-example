package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	pb "github.com/binary-h0/grpc-all-example/ex04/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func uploadFile(client pb.FileServiceClient, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Failed to open file: %v", err)
	}
	defer file.Close()

	stream, err := client.UploadFile(context.Background())
	if err != nil {
		log.Fatalf("Failed to upload file: %v", err)
	}

	buffer := make([]byte, 1024) // 1 KB buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error reading file: %v", err)
		}

		if err := stream.Send(&pb.FileChunk{
			Filename: filePath,
			Content:  buffer[:n],
		}); err != nil {
			log.Fatalf("Error sending chunk: %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Failed to complete upload: %v", err)
	}

	fmt.Println("Upload Response:", res.Message)
}

func downloadFile(client pb.FileServiceClient, filename, outputPath string) {
	stream, err := client.DownloadFile(context.Background(), &pb.FileRequest{Filename: filename})
	if err != nil {
		log.Fatalf("Failed to download file: %v", err)
	}

	file, err := os.Create(outputPath)
	if err != nil {
		log.Fatalf("Failed to create output file: %v", err)
	}
	defer file.Close()

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Error receiving chunk: %v", err)
		}

		_, err = file.Write(chunk.Content)
		if err != nil {
			log.Fatalf("Error writing to file: %v", err)
		}
	}

	fmt.Println("Download complete:", outputPath)
}

func main() {
	conn, err := grpc.NewClient("localhost:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewFileServiceClient(conn)

	// Upload a file
	uploadFile(client, "storage/client_storage/test_upload.txt")

	// Download the file
	downloadFile(client, "test_upload.txt", "storage/client_storage/test_download.txt")
}
