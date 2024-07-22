package main

import (
	"os"
	"testing"
)

func TestNewDeck(t *testing.T) {
	d := newDeck()

	if len(d) != 52 {
		t.Errorf("Expected length of deck: 52, but got %v", len(d))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card: \"Ace of Hearts\", but got %v", d[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card: \"King of Clubs\", but got %v", d[len(d)-1])
	}
}

func TestSaveToFileAndNewDeckFromFile(t *testing.T) {
	os.Remove("./output/_decktesting.txt")

	d := newDeck()
	d.saveToFile("_decktesting")

	newD := newDeckFromFile("_decktesting")

	if !compareSlices(d, newD) {
		t.Errorf("The loaded deck is not the same as the saved deck.")
	}

	if len(newD) != 52 {
		t.Errorf("Expected length of  new deck: 52, but got %v", len(newD))
	}

	if d[0] != "Ace of Hearts" {
		t.Errorf("Expected first card of new deck: \"Ace of Hearts\", but got %v", newD[0])
	}

	if d[len(d)-1] != "King of Clubs" {
		t.Errorf("Expected first card of new deck: \"King of Clubs\", but got %v", newD[len(d)-1])
	}

	os.Remove("./output/_decktesting.txt")
}

func compareSlices(a, b deck) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
