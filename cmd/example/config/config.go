package config

import (
	"log"
	"os"
)

var Token string
var BotName string

func Init() {
	token, present := os.LookupEnv("TOKEN")
	if !present {
		log.Fatalln("No token found")
	} else {
		Token = token
	}

	botName, present := os.LookupEnv("BOT_NAME")
	if !present {
		log.Fatalln("No bot name found")
	} else {
		BotName = botName
	}
}
