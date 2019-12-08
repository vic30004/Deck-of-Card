package main

func main() {
	cards := newDeck()
	hand,remainigDeck:=deal(cards,5)
	hand.print()
	remainigDeck.print()
}

