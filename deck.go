package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	    // Open a file for writing. If the file doesn't exist, it will be created.
	f, err := os.Create(filename)
    if err!= nil {
        return err
    }
    defer f.Close()// Ensure the file is closed when the function returns

    // Write data to the file
    _, err = f.WriteString(d.toString())
    if err!= nil {
        return err
    }
    return nil
}

func newDeckFromFile(filename string) (deck, error) {
	var cards deck
    f, err := os.Open(filename)
    if err!= nil {
		fmt.Println(err)
		os.Exit(1)
        // return newDeck(), err
    }
    defer f.Close()
    scanner := bufio.NewScanner(f)
    for scanner.Scan() {
        line := scanner.Text()
        // Split the line by commas
        s := strings.Split(line, ",")
		
        // Append each term to the slice
        for _, term := range s {
            cards = append(cards, term)
        }
    }
    return cards, nil
}

func (d deck) shuffleCards() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d{
		j := r.Intn(len(d)-1)
        d[i], d[j] = d[j], d[i]
	}

}