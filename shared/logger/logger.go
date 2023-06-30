package logger

import (
	"log"
	"os"
)

var Info *log.Logger
var Warn *log.Logger
var Error *log.Logger

func Init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	Info = log.New(os.Stderr, "INFO\t", log.Ldate|log.Ltime|log.Lshortfile)
	Warn = log.New(os.Stderr, "WARN\t", log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
}
