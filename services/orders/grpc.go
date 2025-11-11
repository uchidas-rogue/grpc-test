package main

import (
	"log"
	"net"

	handler "github.com/uchidas-rogue/kitchen/services/orders/handler/orders"
	"github.com/uchidas-rogue/kitchen/services/orders/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type gRPCServer struct {
	addr string
}

func NewGRPCServer(addr string) *gRPCServer {
	return &gRPCServer{addr: addr}
}

func (s *gRPCServer) Run() error {

	// 証明書と秘密鍵の読み込み
	creds, err := credentials.NewServerTLSFromFile("creds/server.crt", "creds/server.key")
	if err != nil {
		log.Fatalf("failed to load TLS keys: %v", err)
	}

	gRPCServer := grpc.NewServer(grpc.Creds(creds))

	// register services here
	orderService := service.NewOrderService()
	handler.NewGrpcOrdersService(gRPCServer, orderService)

	listener, err := net.Listen("tcp", s.addr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting gRPC server on", s.addr)
	return gRPCServer.Serve(listener)
}
