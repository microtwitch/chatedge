package server

import (
	"context"
	"fmt"
	"log"

	"github.com/microtwitch/chatedge/protos"
)

type receiverServer struct {
	protos.UnimplementedEdgeReceiverServer
}

func NewServer() *receiverServer {
	s := &receiverServer{}
	return s
}

func (s *receiverServer) Send(ctx context.Context, chatMessage *protos.ChatMessage) (*protos.Empty, error) {
	log.Println(fmt.Sprintf("#%s %s: %s", chatMessage.Channel, chatMessage.User, chatMessage.Message))
	return &protos.Empty{}, nil
}
