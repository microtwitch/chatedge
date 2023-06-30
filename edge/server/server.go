package server

import (
	"context"

	"github.com/microtwitch/chatedge/edge/reader"
	"github.com/microtwitch/chatedge/protos"
	"github.com/microtwitch/chatedge/shared/logger"
)

type chatEdgeServer struct {
	protos.UnimplementedChatEdgeServer

	reader *reader.Reader
}

func NewServer() *chatEdgeServer {
	s := &chatEdgeServer{reader: reader.NewReader()}
	return s
}

func (s *chatEdgeServer) JoinChat(ctx context.Context, joinRequest *protos.JoinRequest) (*protos.JoinResponse, error) {
	err := s.reader.Join(joinRequest.Channel, joinRequest.Callback)
	if err != nil {
		return nil, err
	}

	return &protos.JoinResponse{Id: "1"}, nil
}

func (s *chatEdgeServer) Read() {
	err := s.reader.Read()
	if err != nil {
		logger.Error.Fatalln(err)
	}
}
