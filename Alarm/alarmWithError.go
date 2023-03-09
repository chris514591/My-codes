package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

func main() {
	givenAlarmLength := os.Args
	var alarmLENGTH int64 = 10

	switch len(givenAlarmLength) {
	case 2:
		alarmLENGTH, _ = strconv.ParseInt(os.Args[1], 10, 64)
		if alarmLENGTH <= 0 {
			file, err := os.OpenFile("alarmWithErrorLog.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
			if err != nil {
				log.Fatal(err)
				log.SetOutput(file)
			}
		}
	default:
		fmt.Println("Alarm used 10 times by default. Please assign alarm count.")
	}

	var i int64 = 1

	fmt.Println("Hello world, Alarm 1!")

	for {
		if i == alarmLENGTH {
			break
		}
		time.Sleep(100 * time.Millisecond)
		i++
		fmt.Println("Alarm", i)
	}
}
