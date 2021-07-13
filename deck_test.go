package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := addDeck()
	if len(d) != 9 {
		t.Errorf("Expected deck length of 9 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])

	}
}

func TestSaveFile(t *testing.T) {
	os.Remove("_decktesting")
	deck := addDeck()
	deck.writeToFile("_decktesting")
	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 9 {
		t.Errorf("Expected deck of cards is 9, but got %v", len(loadedDeck))
	}
	os.Remove("_decktesting")
}
