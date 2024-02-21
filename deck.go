package main

import "fmt"

//  Create a new type of "deck"
//	which is an slice of strings
type deck []string

func (d deck) print() {
	for i, card := range d {
	    fmt.Printf("%d: %s\n", i+1, card)
	}
}

func newDeck() deck {
	var cards deck
	cardValues:= []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}
	cardSuites := []string{"Clubs", "Diamonds", "Hearts", "Spades"}
	for _, suite := range cardSuites {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suite)
		}
	}
	return cards
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}