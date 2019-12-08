package main

// Create A new type of deck
// Assing it a slice of type string
import (
	"fmt"
	"strings"
)

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, cardS := range cardSuits {
		for _, cardV := range cardValues {
			cards = append(cards, cardV+" of "+cardS)

		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}

}

func deal(d deck, handSize int) (deck, deck) {
	return d[0:handSize], d[handSize:]

}

func (d deck) makeString() string {
	return strings.Join([]string(d), ", ")

}
