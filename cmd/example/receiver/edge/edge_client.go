package edge

import (
	"context"

	"github.com/microtwitch/chatedge/protos"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type ChatEdgeClient struct {
	client protos.ChatEdgeClient
}

func NewChatEdgeClient(target string) (*ChatEdgeClient, error) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	conn, err := grpc.Dial(target, opts...)
	if err != nil {
		return nil, err
	}

	client := protos.NewChatEdgeClient(conn)

	return &ChatEdgeClient{client}, nil
}

func (c *ChatEdgeClient) JoinChat(ctx context.Context, channel string, callback string) error {
	joinRequest := protos.JoinRequest{Channel: channel, Callback: callback}
	_, err := c.client.JoinChat(ctx, &joinRequest)
	if err != nil {
		return err
	}

	// TODO: do something with the id in resp

	return nil
}

func (c *ChatEdgeClient) Send(ctx context.Context, token string, user string, channel string, msg string) error {
	sendRequest := protos.SendRequest{Token: token, User: user, Channel: channel, Msg: msg}
	_, err := c.client.Send(ctx, &sendRequest)
	return err
}
