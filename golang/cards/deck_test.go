package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card of Ace of Spades, but got %v", d[0])
	}
	if len(d) != 16 {
		t.Errorf("Expected deck length of 16, but got %v", len(d))
	}
}

func TestSaveToDeckAndNewDeckFromFile(t *testing.T) {
	// Clean up
	os.Remove("_decktesting")
	// Create a temporary file for testing
	// tmpFile, err := os.CreateTemp("", "_decktesting")
	// if err != nil {
	// 	t.Fatalf("Failed to create temporary file: %v", err)
	// }
	// // defer os.Remove(tmpFile.Name())

	// Create a new deck and save it to the temporary file
	deck := newDeck()
	deck.saveToFile("_decktesting")
	// Load the deck from the temporary file
	loadedDeck := newDeckFromFile("_decktesting")
	// Check if the loaded deck matches the original deck
	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected deck length of %v, but got %v", len(deck), len(loadedDeck))
	}
	for i, card := range deck {
		if card != loadedDeck[i] {
			t.Errorf("Expected card %v at index %v, but got %v", card, i, loadedDeck[i])
		}
	}
}
