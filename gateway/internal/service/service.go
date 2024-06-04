package service

import (
	"API/gateway/internal/service/client"
	proto "API/gateway/pkg/gRPC/api"
	"context"
)

type Service interface {
	GetResponse(ctx context.Context) (string, error)
}

type service struct {
	proto.ServiceAClient
}

func NewService1() Service {
	return &service{
		client.New("service1:50051"),
	}

}
func NewService2() Service {
	return &service{
		client.New("service2:50051"),
	}

}
func (s *service) GetResponse(ctx context.Context) (string, error) {
	req := proto.Request{
		Message: "",
	}

	res, err := s.ServiceAClient.GetResponse(ctx, &req)
	return res.Message, err

}
