package main

import (
	"context"
	"net"

	"github.com/microtwitch/chatedge/edge/config"
	"github.com/microtwitch/chatedge/edge/server"
	"github.com/microtwitch/chatedge/protos"
	"github.com/microtwitch/chatedge/receiver"
	"github.com/microtwitch/chatedge/shared/logger"
	"google.golang.org/grpc"
)

func main() {
	logger.Init()
	config.Init()

	logger.Info.Println("Starting server...")

	lis, err := net.Listen("tcp", "localhost:"+config.Port)
	if err != nil {
		logger.Error.Fatalln(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	server := server.NewServer()

	go server.Read()

	protos.RegisterChatEdgeServer(grpcServer, server)

	go grpcServer.Serve(lis)

	runClient()
}

func runClient() {
	logger.Info.Println("Starting client...")

	client, err := receiver.NewChatEdgeClient()
	if err != nil {
		logger.Error.Fatalln(err)
	}

	err = client.JoinChat(context.Background(), "quin69")
	if err != nil {
		logger.Error.Fatalln(err)
	}

	for {
		select {}
	}
}
