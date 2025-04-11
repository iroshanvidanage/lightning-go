package main

import "fmt"

// We can initialize a variable outside of a function,
// but cannot assign any value ( := ) cannot be used outside of a function to initialize a variable
var deckSize int

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	cards := newDeck()

	// declaring only one variable will only give the index of the slice
	// for card := range cards {
	// 	fmt.Println(card)
	// }

	// cards.print()

	// fmt.Println(cards[0:5])

	// hand1, remainDeck := deal(cards, 5)
	// hand2, remainDeck := deal(remainDeck, 5)
	// hand3, remainDeck := deal(remainDeck, 5)
	// hand4, remainDeck := deal(remainDeck, 5)

	// hand1.print()
	// hand3.print()
	// hand2.print()
	// hand4.print()

	fmt.Println(cards.saveToFile("card_deck.txt"))

}
