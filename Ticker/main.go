package main

import (
	"fmt"
	"time"
)

func main() {
	topTicker := time.NewTicker(time.Minute)
	for {
		fmt.Println("Main Task")
		insideTask()
		<-topTicker.C
	}
}

func insideTask() {
	insideTicker := time.NewTicker(time.Second * 10)
	for {
		fmt.Println("inside task")
		<-insideTicker.C
	}
}
