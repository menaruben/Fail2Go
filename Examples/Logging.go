package main

import (
	"log"
	"os"
)

func main() {
	// If the file doesn't exist, create it or append to the file
	file, err := os.OpenFile("MyLogs.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	log.SetOutput(file)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Hello world!")
}
