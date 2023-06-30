package main

import (
	"context"

	"github.com/microtwitch/chatedge/client"
	"github.com/microtwitch/chatedge/shared/logger"
)

func main() {
	logger.Init()

	logger.Info.Println("Starting client...")

	client, err := client.NewChatEdgeClient()
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
