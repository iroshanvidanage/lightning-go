package main

import (
	"fmt"
	"os"
	"strings"
)

// Create a new type of 'deck'
// which is a slice of strings

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	cardValues := []string{"A", "K", "Q", "J", "10", "9", "8", "7", "6", "5", "4", "3", "2"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+"_"+suit)
		}
	}

	return cards

}

// a function used to print the elements within a slice (receiver)
// (d deck) == d is an argument which is passed as an instance
// the print method will be a property of the deck type

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(filename, []byte(d.toString()), 0666)
}
