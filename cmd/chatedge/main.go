package main

import (
	"net"

	"github.com/microtwitch/chatedge/config"
	"github.com/microtwitch/chatedge/logger"
	"github.com/microtwitch/chatedge/protos"
	"github.com/microtwitch/chatedge/server"
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

	grpcServer.Serve(lis)
}
