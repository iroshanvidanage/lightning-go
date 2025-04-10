package main

import "fmt"

// We can initialize a variable outside of a function,
// but cannot assign any value ( := ) cannot be used outside of a function to initialize a variable
var deckSize int

func main() {
	// var card string = "Ace of Spades"
	// following is used only for declaring new variables
	cards := "Ace of Spades"
	cards = "Five of Hearts"

	// var deckSize int
	// Any variable should be initialized
	deckSize = 52

	fmt.Println(cards, deckSize)

	card := newCard()

	fmt.Println(card)

	// This is also valid
	fmt.Println(newCard())

	// Files in the same package can freely call functions in other files
	// the command should contain all the files
	newDeck()

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

// defining a new function
// the return type must be annotated
func newCard() string {
	return "Five of Diamonds"
}
