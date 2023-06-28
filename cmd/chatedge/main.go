package main

import (
	"net"

	"github.com/microtwitch/chatedge/logger"
	"github.com/microtwitch/chatedge/protos"
	"github.com/microtwitch/chatedge/server"
	"google.golang.org/grpc"
)

func main() {
	logger.Init()

	logger.Info.Println("Starting server...")

	lis, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		logger.Error.Fatalln(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	protos.RegisterChatEdgeServer(grpcServer, server.NewServer())

	grpcServer.Serve(lis)
}
