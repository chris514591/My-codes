package main

import (
	"fmt"
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
	fmt.Println("Welcome to Fonteyn Holidayparks!")

	welcomeOrNot()
}
