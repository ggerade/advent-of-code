package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_readInstructions(t *testing.T) {
	got := readInstructions("inputs/exampleInstructions.txt")
	// Sample data from challenge
	want := []Instruction{
		{ move: 1, from: 2, to: 1 },
		{ move: 3, from: 1, to: 3 },
		{ move: 2, from: 2, to: 1 },
		{ move: 1, from: 1, to: 2 },
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	} else {
		fmt.Printf("exampleInstructions: %v\n", got)
	}
}
