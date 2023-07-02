package server

import (
	"context"
	"fmt"

	"github.com/microtwitch/chatedge/protos"
	"github.com/microtwitch/chatedge/shared/logger"
)

type receiverServer struct {
	protos.UnimplementedEdgeReceiverServer
}

func NewServer() *receiverServer {
	s := &receiverServer{}
	return s
}

func (s *receiverServer) Send(ctx context.Context, chatMessage *protos.ChatMessage) (*protos.Empty, error) {
	logger.Info.Println(fmt.Sprintf("#%s %s: %s", chatMessage.Channel, chatMessage.User, chatMessage.Message))
	return &protos.Empty{}, nil
}
