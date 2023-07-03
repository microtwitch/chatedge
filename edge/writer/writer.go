package writer

import (
	"log"
	"time"

	"github.com/microtwitch/chatedge/protos"
)

type Writer struct {
	clients map[string]*Client
	lastMsg map[string]time.Time
}

func NewWriter() *Writer {
	w := Writer{
		clients: make(map[string]*Client),
		lastMsg: make(map[string]time.Time),
	}

	go w.cleanUpLoop()

	return &w
}

func (w *Writer) Say(sendRequest *protos.SendRequest) error {
	client, exists := w.clients[sendRequest.Token]
	if exists {
		err := client.say(sendRequest.Channel, sendRequest.Msg)
		if err != nil {
			return err
		}

		w.lastMsg[sendRequest.Token] = time.Now()

	} else {
		log.Println("Creating new client for user", sendRequest.User)

		client, err := NewClient(sendRequest.User, sendRequest.Token)
		if err != nil {
			return err
		}
		w.clients[sendRequest.Token] = client

		err = client.say(sendRequest.Channel, sendRequest.Msg)
		if err != nil {
			return err
		}

		w.lastMsg[sendRequest.Token] = time.Now()
	}

	return nil
}

func (w *Writer) cleanUpLoop() {
	for {
		for token, timestamp := range w.lastMsg {
			if time.Since(timestamp).Hours() > 24 {
				log.Println("Cleaning up client because of inactivity")
				delete(w.clients, token)
				delete(w.lastMsg, token)
			}
		}

		time.Sleep(1 * time.Minute)
	}
}
