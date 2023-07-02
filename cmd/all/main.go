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

const EDGE_TARGET string = "127.0.0.1:8080"
const RECEIVER_TARGET string = "127.0.0.1:9090"

func main() {
	logger.Init()
	config.Init()

	logger.Info.Println("Starting server on port 8080")

	lis, err := net.Listen("tcp", EDGE_TARGET)
	if err != nil {
		logger.Error.Fatalln(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	server := server.NewServer()

	go server.Read()

	protos.RegisterChatEdgeServer(grpcServer, server)

	go grpcServer.Serve(lis)

	runReceiver()
}

func runReceiver() {
	logger.Info.Println("Starting receiver server on port 9090")

	lis, err := net.Listen("tcp", RECEIVER_TARGET)
	if err != nil {
		logger.Error.Fatalln(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := receiver.NewServer()

	protos.RegisterEdgeReceiverServer(grpcServer, server)

	go grpcServer.Serve(lis)

	client, err := receiver.NewChatEdgeClient(EDGE_TARGET)
	if err != nil {
		logger.Error.Fatalln(err)
	}

	err = client.JoinChat(context.Background(), "tmiloadtesting2", RECEIVER_TARGET)
	if err != nil {
		logger.Error.Fatalln(err)
	}

	for {
		select {}
	}
}
