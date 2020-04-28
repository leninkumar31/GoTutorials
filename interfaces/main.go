package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}

type spanishBot struct{}

func main() {
	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printGreeting(sb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	//custom logic fo generating  english greeting
	return "Hello There!!"
}

func (spanishBot) getGreeting() string {
	//custom logic fo generating  spanish greeting
	return "Hola!"
}
