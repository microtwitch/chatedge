package server

import (
	"context"

	"github.com/microtwitch/chatedge/logger"
	"github.com/microtwitch/chatedge/protos"
)

type chatEdgeServer struct {
	protos.UnimplementedChatEdgeServer
}

func (s *chatEdgeServer) JoinChat(ctx context.Context, joinRequest *protos.JoinRequest) (*protos.JoinResponse, error) {
	logger.Info.Println(joinRequest.Channel)

	return &protos.JoinResponse{Id: "1"}, nil
}

func NewServer() *chatEdgeServer {
	s := &chatEdgeServer{}
	return s
}
