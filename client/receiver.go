package client

import (
	"context"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/microtwitch/chatedge/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ReceiverClient struct {
	client protos.EdgeReceiverClient
}

func NewReceiverClient(callback string) (*ReceiverClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(callback, opts...)
	if err != nil {
		return nil, err
	}

	client := protos.NewEdgeReceiverClient(conn)

	return &ReceiverClient{client}, nil
}

func (c *ReceiverClient) Send(ctx context.Context, msg twitch.PrivateMessage) error {
	chatMessage := protos.ChatMessage{
		Channel: msg.Channel,
		User:    msg.User.Name,
		Message: msg.Message,
	}
	_, err := c.client.Send(ctx, &chatMessage)

	return err
}
