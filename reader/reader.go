package reader

import (
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/microtwitch/chatedge/client"
	"github.com/microtwitch/chatedge/logger"
	"github.com/microtwitch/chatedge/util"
)

// TODO: should include a client, which is used to send the msgs to the receiver
type Receiver struct {
	channels []string
	client   *client.ReceiverClient
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
	receiver, exists := r.receivers[callback]
	if !exists {
		logger.Info.Println("Registering new receiver for callback" + callback)
		client, err := client.NewReceiverClient(callback)
		if err != nil {
			return err
		}

		r.receivers[callback] = &Receiver{
			channels: []string{channel},
			client:   client,
		}
	} else {
		receiver.channels = append(receiver.channels, channel)
	}

	logger.Info.Println("Joining #" + channel)
	r.client.Join(channel)

	return nil
}

func (r *Reader) onPrivateMessage(msg twitch.PrivateMessage) {
	receiverFound := r.distributeMessage(msg)
	if !receiverFound {
		logger.Warn.Println("No receiver found for channel #" + msg.Channel)
	}
}

func (r *Reader) distributeMessage(msg twitch.PrivateMessage) bool {
	receiverFound := false
	for _, receiver := range r.receivers {
		if util.Contains(receiver.channels, msg.Channel) {
			// TODO: send to client
			receiverFound = true
		}
	}

	return receiverFound
}
