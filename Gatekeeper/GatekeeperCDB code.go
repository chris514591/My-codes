package main

import (
	"bufio"
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

	fmt.Println("Welcome to Fonteyn Holidayparks!")
	fmt.Println("Please input your license plate number to gain access to the parking lot.")

	for tries := 0; tries < 5; tries++ {
		// Loop so you have 5 tries

		scanner2 := bufio.NewScanner(os.Stdin)
		scanner2.Scan()
		text := scanner2.Text()
		// Used to scan and use input from terminal

		if text == "test" {
			welcomeOrNot()
			// Use function if input matches license plate number

		} else {
			if tries < 4 {
				fmt.Println("License plate number not found, try again.")
				// With this statement you won't get "License plate... try again"
			}
		}
		if tries == 4 {
			fmt.Println("Out of tries. Try again later")
		}
	}
}
