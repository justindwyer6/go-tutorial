package main

import "fmt"

type bot interface {
	getGreeting() string
}

type englishBot struct{}
type polishBot struct{}

func main() {
	eb := englishBot{}
	pb := polishBot{}

	printGreeting(eb)
	printGreeting(pb)
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func (englishBot) getGreeting() string {
	// VERY custom logic for generating an English greeting
	return "Hey!"
}

func (polishBot) getGreeting() string {
	// VERY custom logic for generating an Polish greeting
	return "Cześć!"
}
