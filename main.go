package main

import "fmt"

func main() {
	cards := deck{newCard(),newCard()}
	cards = append(cards, newCard())
	cards.print()
	fmt.Printf("%T", cards)
}

func newCard() string {

	return "Five of Diamonds"
}
