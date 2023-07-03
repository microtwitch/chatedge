package server

import (
	"context"
	"log"
	"time"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/microtwitch/chatedge/edge/reader"
	"github.com/microtwitch/chatedge/protos"
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
		log.Fatalln(err)
	}
}

func (s *chatEdgeServer) Send(ctx context.Context, sendRequest *protos.SendRequest) (*protos.Empty, error) {
	err := sendToChat(sendRequest)

	return &protos.Empty{}, err
}

func sendToChat(sendRequest *protos.SendRequest) error {
	client := twitch.NewClient(sendRequest.User, sendRequest.Token)

	client.OnConnect(func() {
		client.Say(sendRequest.Channel, sendRequest.Msg)
		time.Sleep(5 * time.Second)
		client.Disconnect()
	})

	err := client.Connect()

	switch err {
	case twitch.ErrClientDisconnected:
		return nil
	default:
		return err
	}
}
