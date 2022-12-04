package main

import (
	"testing"
)

func Test_determineFullyContainedPairs(t *testing.T) {
	t.Run("with example inputs", func(t *testing.T) {
		exampleAssignments := readAssignments("inputs/example.txt")
		got := determineFullyContainedPairs(exampleAssignments)
		want := 2
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("with challenge inputs", func(t *testing.T) {
		challengeAssignments := readAssignments("inputs/challenge.txt")
		got := determineFullyContainedPairs(challengeAssignments)
		want := 595
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}

func Test_determineOverlappedPairs(t *testing.T) {
	t.Run("with example inputs", func(t *testing.T) {
		exampleAssignments := readAssignments("inputs/example.txt")
		got := determineOverlappedPairs(exampleAssignments)
		want := 4
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})

	t.Run("with challenge inputs", func(t *testing.T) {
		challengeAssignments := readAssignments("inputs/challenge.txt")
		got := determineOverlappedPairs(challengeAssignments)
		want := 952
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
