package main

import (
	"testing"
)

func TestDetermineMostCalories(t *testing.T) {
	// Sample data from challenge 
	elveCalories := [][]int {  
		{ 1000, 2000, 3000 },
		{	4000 },
		{ 5000, 6000},
		{	7000,	8000,	9000 },
		{	10000 },
 	}
	got := determineMostCalories(elveCalories)
	want := 24000

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDetermineMostCaloriesFromFile(t *testing.T) {
	// Sample data from challenge 
	elveCalories := fileTo2DimensionalSliceInts("inputs/example.txt");

	got := determineMostCalories(elveCalories)
	want := 24000

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDetermineMostCaloriesChallenge(t *testing.T) {
	// Sample data from challenge 
	elveCalories := fileTo2DimensionalSliceInts("inputs/challenge.txt");

	got := determineMostCalories(elveCalories)
	want := 70720

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDetermineMostCaloriesChallenge2(t *testing.T) {
	// Sample data from challenge 
	elveCalories := fileTo2DimensionalSliceInts("inputs/challenge.txt");

	got := determineTotalOf3MostCalories(elveCalories)
	want := 207148

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
