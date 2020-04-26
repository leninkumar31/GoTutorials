package main

import "fmt"

func main(){
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)
	t := s[:0]
	printSlice(t)
	u := s[2:]
	printSlice(u)

	var v []int
	printSlice(v)
	if v==nil {
		fmt.Println("nil!")
	}
}

func printSlice(s []int){
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}