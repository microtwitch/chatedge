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
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
	config.Init()

	log.Println("Starting server on", config.Address)

	lis, err := net.Listen("tcp", config.Address)
	if err != nil {
		log.Fatalln(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	server := server.NewServer()

	go server.Read()

	protos.RegisterChatEdgeServer(grpcServer, server)

	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatalln(err)
	}
}
