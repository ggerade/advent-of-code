package main

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_readStacks(t *testing.T) {
	got := readStacks("inputs/exampleStacks.txt")
	// Sample data from challenge
	want := []Stack{
		{"Z", "N"},
		{"M", "C", "D"},
		{"P"},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	} else {
		fmt.Printf("exampleStacks: %v\n", got)
	}
}
