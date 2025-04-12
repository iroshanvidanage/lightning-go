package main

import (
	"errors"
	"io/fs"
	"os"
	"testing"
)

var filename string = "test_file"

func TestNewDeck(t *testing.T) {
	d := newDeck()
	len_d := len(d)
	if len_d != 52 {
		t.Errorf("Expected deck length of 52, but got %v", len_d)
	}
}

func TestSaveToFile(t *testing.T) {
	d := newDeck()
	d.saveToFile(filename)

	if _, err := os.Open(filename); err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			t.Errorf("File does not get saved at the default location, error: %v", err)
		}
		if errors.Is(err, fs.ErrPermission) {
			t.Errorf("File doesn't have enough permission to read. %v", err)
		}
	}

	os.Remove(filename)
}

func TestNewDeckFromFile(t *testing.T) {
	d := newDeck()
	d.saveToFile(filename)
	last_d := len(d) - 1

	file := newDeckFromFile(filename)

	if file[0] != "A_Spades" || file[last_d] != "2_Clubs" {
		t.Errorf("The data in the file isn't correct")
	}

	os.Remove(filename)
}
