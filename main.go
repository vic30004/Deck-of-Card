package main

import "fmt"




func main() {
	cards := newDeck()

	cards.print()
	fmt.Printf("%T", cards)

}

