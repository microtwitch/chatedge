package reader

import (
	"github.com/gempir/go-twitch-irc/v4"
	"github.com/microtwitch/chatedge/logger"
)

// TODO: should include a client, which is used to send the msgs to the receiver
type Receiver struct {
	channels []string
}

type Reader struct {
	client *twitch.Client
}

func NewReader() *Reader {
	r := Reader{}

	client := twitch.NewAnonymousClient()
	client.OnPrivateMessage(r.onPrivateMessage)

	r.client = client
	return &r
}

func (r *Reader) Read() error {
	err := r.client.Connect()
	if err != nil {
		return err
	}

	return nil
}

func (r *Reader) Join(channel string, callback string) {
}

func (r *Reader) onPrivateMessage(msg twitch.PrivateMessage) {
	logger.Info.Println(msg)
}
