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

}

// defining a new function
// the return type must be annotated
func newCard() string {
	return "Five of Diamonds"
}
