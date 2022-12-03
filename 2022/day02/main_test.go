package main

import (
	"testing"
)

func TestDetermineTotalScore1(t *testing.T) {
	strategyGuide := []Strategy1{
		{opponent: ROCK, you: PAPER},
		{opponent: PAPER, you: ROCK},
		{opponent: SCISSORS, you: SCISSORS},
	}

	got := determineTotalScore1(strategyGuide)
	want := 15

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDetermineTotalScore2(t *testing.T) {
	strategyGuide := []Strategy2{
		{opponent: ROCK, goal: DRAW},
		{opponent: PAPER, goal: LOSS},
		{opponent: SCISSORS, goal: WIN},
	}

	got := determineTotalScore2(strategyGuide)
	want := 12

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
