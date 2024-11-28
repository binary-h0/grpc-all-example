package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"

	"net"

	pb "github.com/binary-h0/grpc-all-example/ex04/proto"
	"google.golang.org/grpc"
)

type fileServer struct {
	pb.UnimplementedFileServiceServer
}

const storagePath = "./storage/server_storage"

// UploadFile handles file upload via streaming.
func (s *fileServer) UploadFile(stream pb.FileService_UploadFileServer) error {
	var filename string
	var file *os.File

	for {
		chunk, err := stream.Recv()
		if err == io.EOF {
			// Close file and send response
			if file != nil {
				file.Close()
			}
			log.Printf("File %s uploaded successfully", filename)
			return stream.SendAndClose(&pb.UploadStatus{
				Success: true,
				Message: fmt.Sprintf("File %s uploaded successfully", filename),
			})
		}
		if err != nil {
			return err
		}

		if filename == "" {
			filename = chunk.Filename
			filename = filepath.Base(filename)
			log.Printf("Uploading file: %s", filename)
			file, err = os.Create(filepath.Join(storagePath, filename))
			if err != nil {
				return err
			}
		}

		_, err = file.Write(chunk.Content)
		if err != nil {
			return err
		}
	}
}

// DownloadFile handles file download via streaming.
func (s *fileServer) DownloadFile(req *pb.FileRequest, stream pb.FileService_DownloadFileServer) error {
	filepath := filepath.Join(storagePath, req.Filename)

	file, err := os.Open(filepath)
	if err != nil {
		return err
	}
	defer file.Close()

	buffer := make([]byte, 1024) // 1 KB buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		if err := stream.Send(&pb.FileChunk{
			Filename: req.Filename,
			Content:  buffer[:n],
		}); err != nil {
			return err
		}
	}

	return nil
}

func main() {
	if err := os.MkdirAll(storagePath, os.ModePerm); err != nil {
		log.Fatalf("Failed to create storage directory: %v", err)
	}

	lis, err := net.Listen("tcp", ":50054")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterFileServiceServer(s, &fileServer{})

	log.Println("File server is running on port 50054...")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
