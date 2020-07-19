package main

import (
	"fmt"
	"sync"
)

func main() {
	tokenBucket := NewTokenBucket(2, 2)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			if tokenBucket.Take() {
				fmt.Println("Available")
			} else {
				fmt.Println("Not Available")
			}
		}(&wg)
	}
	wg.Wait()
}
