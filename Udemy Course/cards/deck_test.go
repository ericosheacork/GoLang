package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length 52, but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expect Ace of Spades but got %v", d[0])
	}

	if d[len(d)-1] != "Queen of Clubs" {
		t.Errorf("Expected Queen of clubs but got %v", d[len(d)-1])
	}

}

func TestSaveAndLoad(t *testing.T) {
	os.Remove("_decktesting")

	deck := newDeck()
	deck.saveDeck("_decktesting")

	loadedDeck := readDeckFromFile("_decktesting")

	if len(loadedDeck) != 52 {
		t.Errorf("Expected 52 cards in deck got %v", len(loadedDeck))
	}

	os.Remove("_decktesting")

}
