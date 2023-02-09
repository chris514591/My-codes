package main

import (
	"fmt"
	"time"
)

//V2

func main() {
	fmt.Println("Welkom bij Fonteyn vakantieparken!!!")
	currentTime := time.Now()
	hour := currentTime.Hour()

	if hour >= 7 && hour < 12 {
		fmt.Println("Goedemorgen!")
	} else if hour >= 12 && hour < 18 {
		fmt.Println("Goedemiddag!")
	} else if hour >= 18 && hour < 23 {
		fmt.Println("Goedenavond!")
	} else if hour >= 23 && hour < 7 {
		fmt.Println("Sorry, de parkeerplaats is 's nachts gesloten!")
	}
}
