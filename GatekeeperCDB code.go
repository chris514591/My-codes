package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func welcomeOrNot() {

	currentTime := time.Now()
	hour := currentTime.Hour()

	if hour >= 7 && hour < 12 {
		fmt.Println("Goodmorning!")
	} else if hour >= 12 && hour < 18 {
		fmt.Println("Good afternoon!")
	} else if hour >= 18 && hour < 23 {
		fmt.Println("Good evening!")
	} else if hour >= 23 && hour < 7 {
		fmt.Println("The parking lot is closed. Sorry for the inconvenience.")
	}
}

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Welcome to Fonteyn Holidayparks!")
	fmt.Println("Please input your license plate number to gain access to the parking lot.")

	text, _ := reader.ReadString('\n')
	// convert CRLF to LF
	text = strings.Replace(text, "\n", "", -1)

	if strings.Compare("test", text) == 0 {
		welcomeOrNot()
	} else {
		fmt.Println("License plate number not found, try again.")
	}
}
