package server

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	Server *grpc.Server
	Port   string
}

func NewGRPCServer(port string) *GRPCServer {
	server := grpc.NewServer()
	return &GRPCServer{
		Server: server,
		Port:   port,
	}
}

func (s *GRPCServer) Start() {
	fmt.Println("Starting gRPC server on port", s.Port)
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", s.Port))
	if err != nil {
		panic(err)
	}
	s.Server.Serve(lis)
}
