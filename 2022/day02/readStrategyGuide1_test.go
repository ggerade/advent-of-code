package main

import (
	"reflect"
	"testing"
)

func TestFileRead(t *testing.T) {
	got := readStrategyGuide1("inputs/example.txt")
	// Sample data from challenge
	want := []Strategy1{
		{opponent: ROCK, you: PAPER},
		{opponent: PAPER, you: ROCK},
		{opponent: SCISSORS, you: SCISSORS},
	}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}
