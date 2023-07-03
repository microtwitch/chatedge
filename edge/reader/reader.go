package reader

import (
	"context"
	"log"

	"github.com/gempir/go-twitch-irc/v4"
	"github.com/microtwitch/chatedge/edge/receiver"
	"github.com/microtwitch/chatedge/shared/util"
)

type Receiver struct {
	channels []string
	client   *receiver.ReceiverClient
}

type Reader struct {
	client    *twitch.Client
	receivers map[string]*Receiver
}

func NewReader() *Reader {
	r := Reader{}

	client := twitch.NewAnonymousClient()
	client.OnPrivateMessage(r.onPrivateMessage)

	r.client = client
	r.receivers = make(map[string]*Receiver)
	return &r
}

func (r *Reader) Read() error {
	err := r.client.Connect()
	if err != nil {
		return err
	}

	return nil
}

func (r *Reader) Join(channel string, callback string) error {
	recv, exists := r.receivers[callback]
	if !exists {
		log.Println("Registering new receiver for callback", callback)
		client, err := receiver.NewReceiverClient(callback)
		if err != nil {
			return err
		}

		r.receivers[callback] = &Receiver{
			channels: []string{channel},
			client:   client,
		}
	} else {
		recv.channels = append(recv.channels, channel)
	}

	log.Println("Joining #" + channel)
	r.client.Join(channel)

	return nil
}

func (r *Reader) onPrivateMessage(msg twitch.PrivateMessage) {
	receiverFound := r.distributeMessage(context.Background(), msg)
	if !receiverFound {
		log.Println("No receiver found for channel #" + msg.Channel)
	}
}

func (r *Reader) distributeMessage(ctx context.Context, msg twitch.PrivateMessage) bool {
	receiverFound := false
	for _, receiver := range r.receivers {
		if util.Contains(receiver.channels, msg.Channel) {
			receiverFound = true
			err := receiver.client.Send(ctx, msg)
			if err != nil {
				// TODO: handle too many errors, kick out receiver
				log.Println(err)
			}
		}
	}

	return receiverFound
}
