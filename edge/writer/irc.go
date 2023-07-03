package writer

import (
	"fmt"
	"net"
)

type Client struct {
	conn net.Conn
}

func NewClient(user string, token string) (*Client, error) {
	connection, err := net.Dial("tcp", "irc.chat.twitch.tv:6667")
	if err != nil {
		return nil, err
	}

	_, err = connection.Write([]byte(fmt.Sprintf("PASS %s\r\n", token)))
	if err != nil {
		return nil, err
	}

	_, err = connection.Write([]byte(fmt.Sprintf("NICK %s\r\n", user)))
	if err != nil {
		return nil, err
	}

	return &Client{conn: connection}, nil
}

func (c *Client) say(channel string, msg string) error {
	_, err := c.conn.Write([]byte(fmt.Sprintf("PRIVMSG #%s :%s\r\n", channel, msg)))
	if err != nil {
		return err
	}

	return nil
}
