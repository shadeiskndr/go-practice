package main

import "fmt"

//Create a new type of 'deck'
//which is a slice of strings
type deck []string

func newDeck() deck {
	cards := deck{}
	cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace", "Two", "Three", "Four"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value+" of "+suit)
		}
	}

	// return cards[:3]
	// return cards[12:]
	// return cards[len(cards)-3:]
	// return cards[3:6]
	return cards
}

func (couldbeAnything deck) print() {
	for i, card := range couldbeAnything {
		fmt.Println(i, card)
	}
}

func deal(d deck, handSize int) (deck, deck){
	return d[:handSize], d[handSize:]
}

