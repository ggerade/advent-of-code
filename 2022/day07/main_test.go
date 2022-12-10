package main

import (
	"testing"
)

func Test_determineStartPacketMarker(t *testing.T) {
	t.Run("with example input #1", func(t *testing.T) {
		got := determine("mjqjpqmgbljsphdztnvjfqwrcgsmlb")	
		want := 7		
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	})
}
