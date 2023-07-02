package main

import (
	"context"

	"github.com/microtwitch/chatedge/receiver/edge"
	"github.com/microtwitch/chatedge/shared/logger"
)

func main() {
	logger.Init()

	logger.Info.Println("Starting client...")

	client, err := edge.NewChatEdgeClient("localhost:8080")
	if err != nil {
		logger.Error.Fatalln(err)
	}

	err = client.JoinChat(context.Background(), "tmiloadtesting2", "localhost:9090")
	if err != nil {
		logger.Error.Fatalln(err)
	}

	for {
		select {}
	}
}
