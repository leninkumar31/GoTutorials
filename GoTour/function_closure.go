package main 

import "fmt"

func addder() func(int) int{
	sum := 0
	return func(x int) int {
		sum += x
		return sum
	}
}

func main(){
	pos, neg := addder(), addder()
	for i:=0 ; i<10; i++ {
		fmt.Println(pos(i),neg(-2*i))
	}
}