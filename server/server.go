package server

import (
	"context"

	"github.com/microtwitch/chatedge/logger"
	"github.com/microtwitch/chatedge/protos"
	"github.com/microtwitch/chatedge/reader"
)

type chatEdgeServer struct {
	protos.UnimplementedChatEdgeServer

	reader *reader.Reader
}

func (s *chatEdgeServer) JoinChat(ctx context.Context, joinRequest *protos.JoinRequest) (*protos.JoinResponse, error) {
	logger.Info.Println(joinRequest.Channel)

	s.reader.Join(joinRequest.Channel, joinRequest.Callback)
	return &protos.JoinResponse{Id: "1"}, nil
}

func NewServer() *chatEdgeServer {
	s := &chatEdgeServer{reader: reader.NewReader()}
	return s
}
