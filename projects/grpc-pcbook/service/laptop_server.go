package service

import (
	"context"
	"log"
	example_pb "pcbook/pb/example.pb"
)

// LaptopServer is the server that provides laptop services
type LaptopServer struct {
}

// NewLaptopServer returns a new LaptopServer
func NewLaptopServer() *LaptopServer {
	return &LaptopServer{}
}

// CreateLaptop is an unary RPC to create a new laptop
func (s *LaptopServer) CreateLaptop(ctx context.Context, req *example_pb.CreateLaptopReq) (*example_pb.CreateLaptopRsp, error) {
	laptop := req.GetLaptop()
	log.Printf("receive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		// check if its valid uuid

	}
	return nil
}
