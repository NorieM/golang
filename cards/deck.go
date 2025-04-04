package main

import "fmt"

// create new type - deck
// slice of strings
type deck []string

func newDeck() deck {
	cards :=deck{}
	cardSuits :=[]string{"Spades", "Diamonds", "Hearts", "Clubs"}
	cardValues := []string{"Ace","Two","Threee","Four","Five","Six","Seven","Eight","Nine","Jack","Queen", "King"}

	for _, suit := range cardSuits {
		for _, value := range cardValues {
			cards = append(cards, value + " of " + suit)
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
	return d[:handSize], d[handSize:]
}