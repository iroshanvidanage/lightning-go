package main

// this script contains any auxiliary code that is written to understand the main code
// functionalities

import "fmt"

func auxiliary() {

	// Creating a new slice
	deckCards := []string{newCard(), "Ace of Diamonds"}

	fmt.Println(deckCards)

	// append will not modify the existing array but will create a new slice
	// and assign it to the variable.
	deckCards = append(deckCards, "Six of Spades")

	fmt.Println(deckCards)

	// For loop
	for i, card := range deckCards {
		fmt.Println(i, card)
	}

	for index := 0; index < 10; index++ {
		if index%2 != 0 {
			continue
		}
		fmt.Println("1st type", index)
	}

	for index := range 10 {
		if index%2 != 0 {
			continue
		}
		fmt.Println("2nd type", index)
	}
}

func newDeck() {
	fmt.Println("The deck has 52 cards")
}
