package main

// this script contains any auxiliary code that is written to understand the main code
// functionalities

import "fmt"

func auxiliary() {
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
