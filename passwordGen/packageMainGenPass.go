package main

import (
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func RNG(min, max int) int {
	return rand.Intn(max-min) + min
	// Function generates random integers that belong to a given range using the rand.Intn(). Source: https://opensource.com/article/18/5/creating-random-secure-passwords-go
}

func main() {
	// Function to generate random numbers that will convert to ASCII characters
	MIN := 0
	MAX := 94
	// Min and max type of unique character
	SEED := time.Now().Unix()
	var LENGTH int64 = 10
	// Specifying default Length

	givenLength := os.Args
	// User can specify password length when executing
	switch len(givenLength) {
	case 2:
		// Case 2, case 1 is the default case
		LENGTH, _ = strconv.ParseInt(os.Args[1], 10, 64)
		if LENGTH <= 0 {
			// <= 0 for no rand.Intn panic
			LENGTH = 10
			// if no Length is given Length = 10
		}
	default:
		fmt.Println("Using 10 characters for default!")
	}

	rand.Seed(SEED)
	// Used to keep generating random integers for characters
	startChar := "!"
	// Starting at ! in the ASCII table to get usable characters
	var i int64 = 1
	for {
		myRand := RNG(MIN, MAX)
		newChar := string(startChar[0] + byte(myRand))
		fmt.Print(newChar)
		// Print password per generated character
		if i == LENGTH {
			break
			// Keeps adding one until the default or given Length is reached
		}
		i++
	}
}
