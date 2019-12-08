package main


func main() {
	cards := newDeck()
	cards.saveToFile("New Deck")
	hand, remainigDeck := deal(cards, 5)
	hand.print()
	remainigDeck.print()
	
}
