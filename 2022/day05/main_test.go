package main

import (
	"testing"
)

func Test_determineTopsOfStacks9000(t *testing.T) {
	t.Run("with example inputs", func(t *testing.T) {
		stacks := readStacks("inputs/exampleStacks.txt")
		moves := readInstructions("inputs/exampleInstructions.txt")
		got := determineTopsOfStacks9000(stacks, moves)
	
		want := "CMZ"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("with challenge inputs", func(t *testing.T) {
		stacks := readStacks("inputs/challengeStacks.txt")
		moves := readInstructions("inputs/challengeInstructions.txt")
	
		got := determineTopsOfStacks9000(stacks, moves)
		want := "FRDSQRRCD"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}

func Test_determineTopsOfStacks9001(t *testing.T) {
	t.Run("with example inputs", func(t *testing.T) {
		stacks := readStacks("inputs/exampleStacks.txt")
		moves := readInstructions("inputs/exampleInstructions.txt")
		got := determineTopsOfStacks9001(stacks, moves)
	
		want := "MCD"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})

	t.Run("with challenge inputs", func(t *testing.T) {
		stacks := readStacks("inputs/challengeStacks.txt")
		moves := readInstructions("inputs/challengeInstructions.txt")
	
		got := determineTopsOfStacks9001(stacks, moves)
		want := "HRFTQVWNN"
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
}
