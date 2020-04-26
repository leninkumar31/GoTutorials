package main

import "fmt"

func main(){
	primes := [6]int{2,3,5,7,11,13}
	fmt.Println(primes)
	var a []int = primes[0:2]
	var b []int = primes[1:3]
	fmt.Println(a,b)

	b[0] = 17
	fmt.Println(a,b)
	fmt.Println(primes)
}