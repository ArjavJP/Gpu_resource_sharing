package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/ArjavJP/Gpu_resource_sharing/proto"
)

func main() {
	// Create a listener on port 50052
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the scheduler service
	yourproto.RegisterSchedulerServiceServer(grpcServer, &SchedulerService{})

	// Start the gRPC server
	fmt.Println("Scheduler service listening on port 50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

type SchedulerService struct{}

func (ss *SchedulerService) AllocateGPU(ctx context.Context, req *yourproto.GPURequest) (*yourproto.GPUResponse, error) {
	// GPU allocation logic
	return &yourproto.GPUResponse{Status: "GPU allocated successfully"}, nil
}
