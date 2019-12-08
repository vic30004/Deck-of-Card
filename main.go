package main

import "fmt"

func main() {
	cards := newDeck()
	hand, remainigDeck := deal(cards, 5)
	hand.print()
	remainigDeck.print()
	fmt.Println(cards.makeString())
}
