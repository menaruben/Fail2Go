package main

import (
	"fail2go/winfw"
	"log"
	"os"
)

func main() {
	// Set up logging
	logFile, err := os.OpenFile("Fail2Go.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Can't open logfile: %v", err)
	}
	defer logFile.Close()

	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds)

	// must run as administrator!
	winfw.Ban_IP("192.168.140.10")
	winfw.Unban_IP("192.168.140.10")
}
