package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()
	expectedLen := 52

	if len(d) != expectedLen {
		t.Errorf("Expected deck length of %v, but got %v", expectedLen, len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected first card to be Ace of Spades, but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected last card to be King of Clubs, but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	fileName := "_deckTesting"
	os.Remove(fileName)

	deck := newDeck()
	deck.saveToFile(fileName)

	loadedDeck := newDeckFromFile(fileName)

	expectedDeckLen := 52

	if len(loadedDeck) != expectedDeckLen {
		t.Errorf("Expected %v cards in deck, got %v", expectedDeckLen, len(loadedDeck))
	}

	os.Remove(fileName)
}
