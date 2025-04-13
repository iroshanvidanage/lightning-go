package main

import "fmt"

type bot interface {
	getGreeting() string
	getCountry() string
}

type englishBot struct{}
type spanishBot struct{}

func interfPlay() {
	fmt.Println("Interfaces")

	eb := englishBot{}
	sb := spanishBot{}

	printGreeting(eb)
	printCountry(eb)

	printGreeting(sb)
	printCountry(sb)

}

func (englishBot) getGreeting() string {
	return "Hi there!"
}

func (spanishBot) getGreeting() string {
	return "Hola!"
}

func (englishBot) getCountry() string {
	return "I'm from England"
}

func (spanishBot) getCountry() string {
	return "I'm from France"
}

func printGreeting(b bot) {
	fmt.Println(b.getGreeting())
}

func printCountry(b bot) {
	fmt.Println(b.getCountry())
}
