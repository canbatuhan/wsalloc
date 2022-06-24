package main

import (
	"context"

	pb "workstation.pb.go"
)

type server struct {
	// Server implementation
}

func (s *server) Connect(ctx context.Context, req *pb.ConnectionReq) (*pb.ConnectionRes, error) {
	// Handles connection request
}

func (s *server) AllocateMachine(ctx context.Context, req *pb.AllocReq) (*pb.AllocRes, error) {
	// Handles machine allocation requests
}

func (s *server) ReleaseMachine(ctx context.Context, req *pb.ReleaseReq) (*pb.ReleaseRes, error) {
	// Handles releases
}
