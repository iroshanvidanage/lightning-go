package main

// We can initialize a variable outside of a function,
// but cannot assign any value ( := ) cannot be used outside of a function to initialize a variable
var deckSize int

func main() {
	cards := deck{"Ace of Diamonds", newCard()}
	cards = append(cards, "Six of Spades")

	// declaring only one variable will only give the index of the slice
	// for card := range cards {
	// 	fmt.Println(card)
	// }

	cards.print()
}

// defining a new function
// the return type must be annotated
func newCard() string {
	return "Five of Diamonds"
}
