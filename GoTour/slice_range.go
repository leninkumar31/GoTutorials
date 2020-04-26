package main 

import "fmt"


func main(){
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1<<i
	}
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}