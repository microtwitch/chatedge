package config

import (
	"os"
)

var clientId string
var secret string
var userID string
var callback string
var eventsubSecret string
var botUser string
var botToken string
var channel string

func ClientId() string {
	return clientId
}

func Secret() string {
	return secret
}

func UserID() string {
	return userID
}

func Callback() string {
	return callback
}

func EventsubSecret() string {
	return eventsubSecret
}

func BotUser() string {
	return botUser
}

func BotToken() string {
	return botToken
}

func Channel() string {
	return channel
}

func Init() {
	clientId = os.Getenv("CLIENT_ID")
	secret = os.Getenv("SECRET")
	userID = os.Getenv("USER_ID")
	callback = os.Getenv("CALLBACK")
	eventsubSecret = os.Getenv("EVENTSUB_SECRET")
	botUser = os.Getenv("BOT_USER")
	botToken = os.Getenv("BOT_TOKEN")
	channel = os.Getenv("CHANNEL")
}
