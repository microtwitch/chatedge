package main

import (
	"context"

	"github.com/microtwitch/chatedge/client"
	"github.com/microtwitch/chatedge/logger"
)

func main() {
	logger.Init()

	logger.Info.Println("Starting client...")

	client, err := client.NewChatEdgeClient()
	if err != nil {
		logger.Error.Fatalln(err)
	}

	err = client.JoinChat(context.Background(), "turtoise")
	if err != nil {
		logger.Error.Fatalln(err)
	}
}
