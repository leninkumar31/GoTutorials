package main

import "fmt"

func main() {
	numbers := []int{}
	for i := 0; i <= 10; i++ {
		numbers = append(numbers, i)
	}

	for _, val := range numbers {
		if val%2 == 0 {
			fmt.Printf("%v is even\n", val)
		} else {
			fmt.Printf("%v is odd\n", val)
		}
	}
}
