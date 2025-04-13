package main

import (
	"fmt"
	"math/rand"
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

func newDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	s := strings.Split(string(bs), ",")
	return deck(s)
}

func (d deck) shuffleDeck() {
	for i := range d {
		rN := rand.Intn(len(d) - 1)
		if i != rN {
			d[i], d[rN] = d[rN], d[i]
		}
	}
}

func deckPlay() {
	// card_deck := deck{"Ace of Diamonds", newCard()}
	// card_deck = append(card_deck, "Six of Spades")

	// new deck of cards
	card_deck := newDeck()

	// declaring only one variable will only give the index of the slice
	// for card := range card_deck {
	// 	fmt.Println(card)
	// }

	// card_deck.print()

	fmt.Println(card_deck[0:5])

	// hand1, remainDeck := deal(card_deck, 5)
	// hand2, remainDeck := deal(remainDeck, 5)
	// hand3, remainDeck := deal(remainDeck, 5)
	// hand4, remainDeck := deal(remainDeck, 5)

	// hand1.print()
	// hand3.print()
	// hand2.print()
	// hand4.print()

	// Save to file
	fmt.Println(card_deck.saveToFile("card_deck.txt"))

	// a new deck from the file
	cards := newDeckFromFile("card_deck.txt")
	fmt.Println("New deck from file")
	cards.print()

	fmt.Println("Shuffle the cards")
	cards.shuffleDeck()
	cards.print()
}
