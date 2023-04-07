package main

import (
	"fail2go/SSHLogParsing"
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
	bannable_hosts := SSHLogParsing.Get_bannable_hosts(2)

	for i := 0; i < len(bannable_hosts); i++ {
		winfw.Ban_IP(bannable_hosts[i])
	}

	/* TODO:
	- add Release date from sshjail
	- save banned hosts to sql or toml ("ip = release date" form maybe)
	- import banned hosts at start */

	// winfw.Unban_IP("")
}
