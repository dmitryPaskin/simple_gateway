package client

import (
	proto "API/gateway/pkg/gRPC/api"
	"context"
	"google.golang.org/grpc"
)

type client struct {
	target string
}

func New(target string) proto.ServiceAClient {
	return &client{target: target}
}

func (c *client) GetResponse(ctx context.Context, in *proto.Request, opts ...grpc.CallOption) (*proto.Response, error) {
	conn, err := grpc.Dial(c.target, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	cl := proto.NewServiceAClient(conn)

	res, err := cl.GetResponse(ctx, in)
	return res, err
}
