package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}

	if d[0] != "Ace of Spades" {
		t.Errorf("Expected `Ace of Spades` but found `%v` at 0th index", d[0])
	}
}

func TestSavetoDeckAndNewDecFromFile(t *testing.T) {
	testingFileName := "_decTesting"
	// remove file if already exist
	os.Remove(testingFileName)

	// Created new Deck
	deck := newDeck()
	// save to file
	deck.saveToFile(testingFileName)

	loadedDeck := getDeckFromFile(testingFileName)

	if len(loadedDeck) != len(deck) {
		t.Errorf("Expected length of %v but found deck of length %v ", len(deck), len(loadedDeck))
	}

	if loadedDeck[0] != deck[0] {
		t.Errorf("First Element did not match with original deck")
	}

	//Cleaning up the created deck
	os.Remove(testingFileName)
}
