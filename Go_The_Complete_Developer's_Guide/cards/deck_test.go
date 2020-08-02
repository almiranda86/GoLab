package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deckLen := 16
	d := newDeck()

	if len(d) != deckLen {
		t.Errorf("Expected deck len of %v, but got %v\n", deckLen, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "Four of Clubs" {
		t.Errorf("Expected last card of Four of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	os.Remove("_decktesting")

	deckLen := 16

	deck := newDeck()
	deck.saveToFile("_decktesting")

	loadDeck := newDeckFromFile("_decktesting")

	if len(loadDeck) != deckLen {
		t.Errorf("Expected deck len of %v, but got %v\n", deckLen, len(loadDeck))
	}

	os.Remove("_decktesting")
}
