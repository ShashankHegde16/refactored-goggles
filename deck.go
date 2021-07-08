package main

import "fmt"

type deck []string

func (n deck) print() {
	for this, card := range n {
		fmt.Println(this, card)
	}
}

func addDeck() deck {
	cards := deck{}

	cardSuites := []string{"Spades", "Diamonds", "Hearts"}
	cardValues := []string{"Ace", "Two", "Three"}

	for _, suit := range cardSuites {
		for _, values := range cardValues {
			cards = append(cards, values+" of "+suit)
		}
	}
	return cards
}
