package main

import (
	"context"
	"log"
	"net"

	"github.com/microtwitch/chatedge/cmd/example/config"
	"github.com/microtwitch/chatedge/cmd/example/receiver/edge"
	receiver_server "github.com/microtwitch/chatedge/cmd/example/receiver/server"
	"github.com/microtwitch/chatedge/edge/server"
	"github.com/microtwitch/chatedge/protos"
	"google.golang.org/grpc"
)

const EDGE_TARGET string = "127.0.0.1:8080"
const RECEIVER_TARGET string = "127.0.0.1:9090"

func main() {
	config.Init()
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	log.Println("Starting server on port 8080")

	lis, err := net.Listen("tcp", EDGE_TARGET)
	if err != nil {
		log.Fatalln(err)
	}

	var opts []grpc.ServerOption

	grpcServer := grpc.NewServer(opts...)
	server := server.NewServer()

	go server.Read()

	protos.RegisterChatEdgeServer(grpcServer, server)

	go func() {
		grpcServer.Serve(lis)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	runReceiver()
}

func runReceiver() {
	log.Println("Starting receiver server on port 9090")

	lis, err := net.Listen("tcp", RECEIVER_TARGET)
	if err != nil {
		log.Fatalln(err)
	}

	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	server := receiver_server.NewServer()

	protos.RegisterEdgeReceiverServer(grpcServer, server)

	go func() {
		err := grpcServer.Serve(lis)
		if err != nil {
			log.Fatalln(err)
		}
	}()

	client, err := edge.NewChatEdgeClient(EDGE_TARGET)
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Send(context.Background(), config.Token, config.BotName, "matthewde", "test1")
	if err != nil {
		log.Fatalln(err)
	}

	err = client.Send(context.Background(), config.Token, config.BotName, "matthewde", "test2")
	if err != nil {
		log.Fatalln(err)
	}

	err = client.JoinChat(context.Background(), "tmiloadtesting2", RECEIVER_TARGET)
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {}
	}
}
