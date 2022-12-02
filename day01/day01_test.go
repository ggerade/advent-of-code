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
