package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	pb "vusb-hello-world/proto"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedVirtualUSBServiceServer
	serverID string
}

func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Hello from: %s, device: %s", req.GetName(), req.GetDeviceId())
	
	message := fmt.Sprintf("Hello %s! Your device %s is connected to server %s", 
		req.GetName(), req.GetDeviceId(), s.serverID)
	
	return &pb.HelloResponse{
		Message:  message,
		ServerId: s.serverID,
	}, nil
}

func main() {
	hostname, _ := os.Hostname()
	serverID := fmt.Sprintf("%s-%d", hostname, os.Getpid())
	
	log.Printf("Starting Virtual USB Hello World Service")
	log.Printf("Server ID: %s", serverID)
	
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}
	
	s := grpc.NewServer()
	pb.RegisterVirtualUSBServiceServer(s, &server{serverID: serverID})
	
	log.Printf("Ready at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
