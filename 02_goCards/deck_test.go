package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	expectedDeckLength := 16
	expectedFirstCard := "Ace of Spades"
	expectedLastCard := "Four of Clubs"

	d := newDeck()
	if len(d) != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckLength, len(d))
	}

	if d[0] != expectedFirstCard {
		t.Errorf("Expected first card to be %v, but got %v", expectedFirstCard, d[0])
	}

	if d[len(d) - 1] != expectedLastCard {
		t.Errorf("Expected first card to be %v, but got %v", expectedLastCard, d[len(d) - 1])
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	filename := "_decktesting"
	expectedDeckLength := 16

	_ = os.Remove(filename)

	deck := newDeck()
	_ = deck.saveToFile(filename)
	loadedDeck := newDeckFromFile(filename)

	if len(loadedDeck) != expectedDeckLength {
		t.Errorf("Expected deck length of %v, but got %v", expectedDeckLength, len(loadedDeck))
	}

	_ = os.Remove(filename)
}