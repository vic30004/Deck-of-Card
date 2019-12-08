package main

// Create A new type of deck
// Assing it a slice of type string
import "fmt"

type deck []string

func (d deck) print(){
	for i, card := range d{
		fmt.Println(i,card)
	}
	
}