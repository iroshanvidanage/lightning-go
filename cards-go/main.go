package main

// We can initialize a variable outside of a function,
// but cannot assign any value ( := ) cannot be used outside of a function to initialize a variable

func main() {
	// cards := deck{"Ace of Diamonds", newCard()}
	// cards = append(cards, "Six of Spades")

	// new deck of cards
	// cards := newDeck()

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

	// Save to file
	// fmt.Println(cards.saveToFile("card_deck.txt"))

	// a new deck from the file
	// cards := newDeckFromFile("card_deck.txt")
	// fmt.Println("New deck from file")
	// cards.print()

	// fmt.Println("Shuffle the cards")
	// cards.shuffleDeck()
	// cards.print()

	int_slice := []int{}
	for i := range 11 {
		int_slice = append(int_slice, i)
	}

	odd_even(int_slice)
}
