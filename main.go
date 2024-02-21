package main

import "fmt"
func main() {
	cards:= newDeck()
	// cards,error :=newDeckFromFile("cards.txt")
	cards.shuffleCards()
	// cards.print()
	fmt.Println(cards.saveToFile("cards.txt"))
	// fmt.Println(cards.toString(), error)
	// hand, remainingCards := deal(cards, 5)
	// hand.print()
	// fmt.Println("--------------------------------")
	// remainingCards.print()
}
