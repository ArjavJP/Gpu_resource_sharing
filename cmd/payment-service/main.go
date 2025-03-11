package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"github.com/ArjavJP/Gpu_resource_sharing/proto"
)

func main() {
	// Set up the gRPC server
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen on port 50051: %v", err)
	}

	// Create a new gRPC server
	grpcServer := grpc.NewServer()

	// Register the payment service
	yourproto.RegisterPaymentServiceServer(grpcServer, &PaymentService{})

	// Start the gRPC server
	fmt.Println("Payment service listening on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

type PaymentService struct{}

func (ps *PaymentService) ProcessPayment(ctx context.Context, req *yourproto.PaymentRequest) (*yourproto.PaymentResponse, error) {
	// Process the payment logic
	return &yourproto.PaymentResponse{Status: "Payment successful"}, nil
}
