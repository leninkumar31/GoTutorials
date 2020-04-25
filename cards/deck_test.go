package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	deck := newDeck()
	if len(deck) != 52 {
		t.Errorf("Expected deck length to be 52. But got %v", len(deck))
	}

	if deck[0] != "Ace of Spades" {
		t.Errorf("Expected Ace of Spades. But got %v", deck[0])
	}

	if deck[51] != "Jack of Clubs" {
		t.Errorf("Expected Jack of Clubs. But got %v", deck[51])
	}
}

func TestSaveToDeckAndTestNewDeckFromFile(t *testing.T) {
	os.Remove("_deckTesting")
	deck := newDeck()
	deck.saveToFile("_deckTesting")
	loadedDeck := newDeckFromFile("_deckTesting")
	if len(loadedDeck) != 52 {
		t.Errorf("Expected loaded deck length to be 52. But got %v", len(loadedDeck))
	}
	os.Remove("_deckTesting")
}
