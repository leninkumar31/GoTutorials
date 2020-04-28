package main

import "fmt"

type shape interface {
	getArea() float64
}

type triangle struct {
	base, height float64
}

type square struct {
	length float64
}

func main() {
	t := triangle{base: 2, height: 1}
	s := square{length: 1}
	printArea(t)
	printArea(s)
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.length * s.length
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}
