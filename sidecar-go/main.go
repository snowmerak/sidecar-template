package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	// "github.com/my-org/core-platform/gen/go/company/auth/v1"
)

func main() {
	port := 50051
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()

	// TODO: Register your services here
	// authv1.RegisterAuthServiceServer(s, &myAuthService{})

	log.Printf("Sidecar Go Server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
