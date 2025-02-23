package main

import (
	"log"
	"net"

	handler "github.com/Sanjaiy/go-grpc/services/orders/handler/orders"
	"github.com/Sanjaiy/go-grpc/services/orders/service"
	"google.golang.org/grpc"
)

type gRPCSever struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCSever {
	return &gRPCSever{addr: addr}
}

func (s *gRPCSever) Run() error {
	lis, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	orderService := service.NewOrderService()
	handler.NewGrpcOrderService(grpcServer, orderService)

	log.Println("Starting gRPC server on", s.addr)

	return grpcServer.Serve(lis)
}