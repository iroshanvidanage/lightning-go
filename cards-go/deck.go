package main

import "fmt"

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

// a function used to print the elements within a slice (receiver)
// (d deck) == d is an argument which is passed as an instance
// the print method will be a property of the deck type

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}
