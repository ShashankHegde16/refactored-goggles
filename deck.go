package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"os"
	"strings"
	"time"
)

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

func deal(d deck, n int) (deck, deck) {
	return d[:n], d[n:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) writeToFile(filename string) error {
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	return err
}

func newDeckFromFile(filename string) deck {
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
	s := strings.Split(string(content), ",")

	return deck(s)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newIndex := r.Intn(len(d) - 1)
		d[i], d[newIndex] = d[newIndex], d[i]
	}
}
