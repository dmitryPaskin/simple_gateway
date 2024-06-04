package server

import (
	"API/service2/internal/service"
	proto "API/service2/pkg/gRPC/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

type Server interface {
	Start()
}

type server struct {
	proto.ServiceAServer
}

func New() Server {
	s := service.New()
	return &server{s}
}

func (s *server) Start() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	ser := grpc.NewServer()
	reflection.Register(ser)
	proto.RegisterServiceAServer(ser, s.ServiceAServer)

	log.Printf("server listening at: %v", ":50051")

	if err = ser.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
