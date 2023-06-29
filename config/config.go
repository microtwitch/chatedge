package config

import (
	"os"
)

var Port string

func Init() {
	port, present := os.LookupEnv("PORT")
	if !present {
		Port = "8080"
	} else {
		Port = port
	}
}
