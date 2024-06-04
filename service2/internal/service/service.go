package service

import (
	proto "API/service2/pkg/gRPC/api"
	"context"
)

type service struct {
	proto.UnimplementedServiceAServer
}

func New() proto.ServiceAServer {
	return &service{}
}

func (s *service) GetResponse(context.Context, *proto.Request) (*proto.Response, error) {
	return &proto.Response{Message: "Response from Service B"}, nil
}
