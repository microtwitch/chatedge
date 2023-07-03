package server

import (
	"context"
	"log"

	"github.com/microtwitch/chatedge/edge/reader"
	"github.com/microtwitch/chatedge/edge/writer"
	"github.com/microtwitch/chatedge/protos"
)

type chatEdgeServer struct {
	protos.UnimplementedChatEdgeServer

	reader *reader.Reader
	writer *writer.Writer
}

func NewServer() *chatEdgeServer {
	s := &chatEdgeServer{reader: reader.NewReader(), writer: writer.NewWriter()}
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
		log.Fatalln(err)
	}
}

func (s *chatEdgeServer) Send(ctx context.Context, sendRequest *protos.SendRequest) (*protos.Empty, error) {
	err := s.writer.Say(sendRequest)
	return &protos.Empty{}, err
}
