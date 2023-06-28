package client

import (
	"context"

	"github.com/microtwitch/chatedge/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ChatEdgeClient struct {
	client protos.ChatEdgeClient
}

func NewChatEdgeClient() (*ChatEdgeClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial("localhost:8080", opts...)
	if err != nil {
		return nil, err
	}

	client := protos.NewChatEdgeClient(conn)

	return &ChatEdgeClient{client}, nil
}

func (c *ChatEdgeClient) JoinChat(ctx context.Context, channel string) error {
	joinRequest := protos.JoinRequest{Channel: channel}
	_, err := c.client.JoinChat(ctx, &joinRequest)

	return err
}
