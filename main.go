package main

import "fmt"
func main() {
	cards:= newDeck()
	cards.shuffleCards()
	// cards.print()
	fmt.Println(cards.saveToFile("cards.txt"))
	// cards,error :=newDeckFromFile("cards.txt")
	// hand, remainingCards := deal(cards, 5)
	
	// hand.print()
	// fmt.Println("--------------------------------")
	// remainingCards.print()
}
