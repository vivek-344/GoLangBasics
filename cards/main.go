package main

import "fmt"

func main() {
	cards := newDeck()

	cards.shuffle()

	hand, remCards := deal(cards, 5)

	hand.print()

	err := remCards.saveToFile("Remaining Cards")

	if err == nil {
		fmt.Println("Saved to File!")
	} else {
		fmt.Println(err)
	}

	newCards := newDeckFromFile("Remaining Cards")

	newCards.print()
}
