package main

import (
	"fmt"
	"os"
	"time"
	// Needed to use certain things in your code, may be added automatically. Assume you need to add them yourself
)

func welcomeOrNot() {

	currentTime := time.Now()
	hour := currentTime.Hour()
	// Translate the current time to currentTime and than to "hour"

	if hour >= 7 && hour < 12 {
		fmt.Println("Goodmorning!")
	} else if hour >= 12 && hour < 18 {
		fmt.Println("Good afternoon!")
	} else if hour >= 18 && hour < 23 {
		fmt.Println("Good evening!")
	} else if hour >= 23 && hour < 7 {
		fmt.Println("The parking lot is closed. Sorry for the inconvenience.")
	}
	// Seperated function to keep de code more clean (My opinion)
}

func main() {
	licensePlateNumber := os.Args[1]
	// Used to insert input right after calling .exe

	fmt.Println("Welcome to Fonteyn Holidayparks!")
	fmt.Println("Please input your license plate number to gain access to the parking lot.")

	if licensePlateNumber == "test" {
		welcomeOrNot()
		// Use function if input matches license plate number
	}
}
