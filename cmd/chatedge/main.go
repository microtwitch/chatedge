package main

import (
	"log"
	"net"

	"github.com/microtwitch/chatedge/edge/config"
	"github.com/microtwitch/chatedge/edge/server"
	"github.com/microtwitch/chatedge/protos"
	"google.golang.org/grpc"
)

func main() {
	config.Init()

	log.Println("Starting server...")

	lis, err := net.Listen("tcp", "127.0.0.1:"+config.Port)
	if err != nil {
		log.Fatalln(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	server := server.NewServer()

	go server.Read()

	protos.RegisterChatEdgeServer(grpcServer, server)

	grpcServer.Serve(lis)
}
