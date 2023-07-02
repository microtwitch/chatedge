package main

import (
	"net"

	"github.com/microtwitch/chatedge/edge/config"
	"github.com/microtwitch/chatedge/edge/server"
	"github.com/microtwitch/chatedge/protos"
	"github.com/microtwitch/chatedge/shared/logger"
	"google.golang.org/grpc"
)

func main() {
	logger.Init()
	config.Init()

	logger.Info.Println("Starting server...")

	lis, err := net.Listen("tcp", "127.0.0.1:"+config.Port)
	if err != nil {
		logger.Error.Fatalln(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	server := server.NewServer()

	go server.Read()

	protos.RegisterChatEdgeServer(grpcServer, server)

	grpcServer.Serve(lis)
}
