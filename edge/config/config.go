package config

import (
	"log"
	"os"
)

var Address string

func Init() {
	address, present := os.LookupEnv("ADDRESS")
	if !present {
		log.Fatalln("No address found")
	} else {
		Address = address
	}
}
