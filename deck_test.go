package main

import (
	"fmt"
	"os"
	"testing"
)
func TestNewDeck(t *testing.T) {
	d := newDeck()
	if len(d)!= 52 {
		t.Errorf("Expected 52 cards, got %d", len(d))
	}

	if d[0]!= "Ace of Clubs" {
        t.Errorf("Expected first card Ace of Clubs, got %s", d[0])
    }

	if d[len(d)-1]!= "King of Spades" {
        t.Errorf("Expected last card King of Spades, got %s", d[len(d)-1])
    }
}

func TestDeal(t *testing.T) {
	cards := newDeck()
    hand, remainingCards := deal(cards, 5)
    hand.print()
    fmt.Println()
    remainingCards.print()
}	

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting.txt")
	deck := newDeck()
    err := deck.saveToFile("_decktesting.txt")
    if err != nil {
		t.Errorf("error saving to file: %v", err)
	}

	loadedDeck, err := newDeckFromFile("_decktesting.txt")
	
    if err != nil {
		t.Errorf("error loading from file to file: %v", err)
	}

	if len(loadedDeck)!= 52 {
		t.Errorf("Expected 52 cards, got %d", len(loadedDeck))
	}

	os.Remove("_decktesting.txt")

}