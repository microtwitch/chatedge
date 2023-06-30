package receiver

import (
	"context"

	"github.com/microtwitch/chatedge/protos"
	"github.com/microtwitch/chatedge/shared/logger"
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
	resp, err := c.client.JoinChat(ctx, &joinRequest)
	if err != nil {
		return err
	}

	logger.Info.Println(resp.Id)

	return err
}