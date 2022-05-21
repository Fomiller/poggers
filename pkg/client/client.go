package client

import (
	"time"

	"github.com/gempir/go-twitch-irc/v3"
	"golang.org/x/oauth2/clientcredentials"
)

var (
	oauth2Config *clientcredentials.Config
)

type PogMessage struct {
	User    string
	Message string
	Time    time.Time
}

// type Reader interface {
// 	OnPrivateMessage()
// }

type Client struct {
	// Reader
	twitch.Client
	client *twitch.Client
	// onPrivateMessage func(message twitch.PrivateMessage) *PogMessage
	// ReturnPrivateMessage func(message twitch.PrivateMessage) *PogMessage
}

// takes callback function that handles the PrivateMessage
// func (c *Client) OnPrivateMessage(callback func(message twitch.PrivateMessage) *PogMessage) {
// 	c.onPrivateMessage = callback
// }

// func (c *Client) ReturnPrivateMessage(message twitch.PrivateMessage) *PogMessage {
// 	// c.ReturnPrivateMessage = callbac
// 	if strings.Contains(strings.ToLower(message.Message), "pog") {
// 		return &PogMessage{message.User.Name, message.Message, message.Time}
// 	} else {
// 		return &PogMessage{}
// 	}
// }

func (c *Client) CreateClient(username, accessToken string) {
	c.client = twitch.NewClient(username, accessToken)
}

func (c *Client) ConnectClient(channel string) {
	c.client.Join(channel)

	err := c.client.Connect()
	if err != nil {
		panic(err)
	}
}
