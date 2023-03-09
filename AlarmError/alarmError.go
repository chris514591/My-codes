package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.OpenFile("alarmErrorLogs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	log.SetOutput(file)

	log.Println("Error 1")
}
