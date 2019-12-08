package main

import "fmt"

func main() {
	cards := newDeck()
	cards.saveToFile("New Deck")
	fmt.Println(newDeckFromFile("New Deck"))
	hand, remainigDeck := deal(cards, 5)
	hand.print()
	remainigDeck.print()

}
