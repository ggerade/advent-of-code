package main

import (
	"reflect"
	"testing"
)

func Test_readAssignmentPairs(t *testing.T) {
	got := readAssignments("inputs/example.txt")
	// Sample data from challenge
	want := []Assignment{
		{begin:2, end:4}, {begin:6, end:8},
		{begin:2, end:3}, {begin:4, end:5},
		{begin:5, end:7}, {begin:7, end:9},
		{begin:2, end:8}, {begin:3, end:7},
		{begin:6, end:6}, {begin:4, end:6},
		{begin:2, end:6}, {begin:4, end:8},
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v want %v", got, want)
	}
}
