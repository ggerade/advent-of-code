package main

import (
	"reflect"
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

func TestFileRead(t *testing.T) {
	got := fileTo2DimensionalSliceInts("example_input.txt");
	// Sample data from challenge 
	want := [][]int {  
		{ 1000, 2000, 3000 },
		{	4000 },
		{ 5000, 6000},
		{	7000,	8000,	9000 },
		{	10000 },
 	}

	 if !reflect.DeepEqual(got, want) {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDetermineMostCaloriesFromFile(t *testing.T) {
	// Sample data from challenge 
	elveCalories := fileTo2DimensionalSliceInts("example_input.txt");

	got := determineMostCalories(elveCalories)
	want := 24000

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDetermineMostCaloriesChallenge(t *testing.T) {
	// Sample data from challenge 
	elveCalories := fileTo2DimensionalSliceInts("challenge.txt");

	got := determineMostCalories(elveCalories)
	want := 70720

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDetermineMostCaloriesChallenge2(t *testing.T) {
	// Sample data from challenge 
	elveCalories := fileTo2DimensionalSliceInts("challenge.txt");

	got := determineTotalOf3MostCalories(elveCalories)
	want := 207148

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
