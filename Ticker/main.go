package main

import (
	"fmt"
	"time"
)

func main() {
	ticker := time.NewTicker(time.Second)

	go func() {
		for range ticker.C {
			fmt.Println("Tick")
		}
	}()
	time.Sleep(5 * time.Second)
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
