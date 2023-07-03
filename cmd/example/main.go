package main

import (
	"context"
	"log"

	"github.com/microtwitch/chatedge/receiver/edge"
)

func main() {
	log.Println("Starting client...")

	client, err := edge.NewChatEdgeClient("localhost:8080")
	if err != nil {
		log.Fatalln(err)
	}

	err = client.JoinChat(context.Background(), "tmiloadtesting2", "localhost:9090")
	if err != nil {
		log.Fatalln(err)
	}

	for {
		select {}
	}
}
