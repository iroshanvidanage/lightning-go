package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type bot interface {
	getGreeting() string
	getCountry() string
}

type logWriter struct{}

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

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("No. of bytes: ", len(bs))

	return len(bs), nil
}

func httpPlay() {
	resp, err := http.Get("https://google.com")
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}

	fmt.Printf("%+v\n", resp)

	// bs := make([]byte, 99999)
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	lw := logWriter{}

	// io.Copy(os.Stdout, resp.Body)
	io.Copy(lw, resp.Body)
}

// #####################################
// Assignement: Interfaces

type shape interface {
	getArea() float64
}

type triangle struct {
	height float64
	base   float64
}

type square struct {
	sideLength float64
}

func (t triangle) getArea() float64 {
	return 0.5 * t.base * t.height
}

func (s square) getArea() float64 {
	return s.sideLength * s.sideLength
}

func printArea(s shape) {
	fmt.Println(s.getArea())
}

func interfaceAssignment() {
	tri := triangle{base: 2, height: 4}
	sq := square{sideLength: 3}

	printArea(tri)
	printArea(sq)
}
