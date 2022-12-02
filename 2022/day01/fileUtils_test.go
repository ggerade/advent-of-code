package main

import (
	"reflect"
	"testing"
)

func TestFileRead(t *testing.T) {
	got := fileTo2DimensionalSliceInts("inputs/example.txt");
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
