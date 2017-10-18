package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected \"Ace of spades\", but got %v", d[0])
	}

	if d[51] != "2 of Clovers" {
		t.Errorf("Expected \"2 of Clovers\", but got %v", d[51])
	}
}
func TestSaveToDeckAndNewDeckTestFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadedDeck := newDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck, get %v", len(loadedDeck))
	}

	// os.Remove("_decktesting")
}
