package main

import (
	"testing"
)

func TestSplitRucksackCompartments(t *testing.T) {
	got1, got2 := splitRucksackCompartments("vJrwpWtwJgWrhcsFMMfFFhFp")
	want1 := "vJrwpWtwJgWr"
	want2 := "hcsFMMfFFhFp"

	if got1 != want1 {
		t.Errorf("got %s want %s", got1, want1)
	}
	if got2 != want2 {
		t.Errorf("got %s want %s", got2, want2)
	}
}

func TestFindDuplicate(t *testing.T) {
	got := findFirstDuplicate("vJrwpWtwJgWr", "hcsFMMfFFhFp")
	want := 'p'

	if got != want {
		t.Errorf("got %c want %c", got, want)
	}
}

func TestDeterminePriorityLower(t *testing.T) {
	got := determinePriority('a')
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDeterminePriorityUpper(t *testing.T) {
	got := determinePriority('A')
	want := 27

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDetermineTotalPriority_Example(t *testing.T) {
	// Sample data from challenge
	rucksacks := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	got := determineTotalPriority(rucksacks)
	want := 157

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestDetermineTotalPriority_Challenge(t *testing.T) {
	rucksacks := readRucksacks("inputs/challenge.txt")
	got := determineTotalPriority(rucksacks)
	want := 8349

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

func TestFindGroupBadge(t *testing.T) {
	rucksacks := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}
	got := findGroupBadge(rucksacks[0], rucksacks[1], rucksacks[2])
	want := 'r'

	if got != want {
		t.Errorf("got %c want %c", got, want)
	}
}

func TestDetermineGroupBadgeTotalPriority_Example(t *testing.T) {
	// Sample data from challenge
	rucksacks := []string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
		"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
		"ttgJtRGJQctTZtZT",
		"CrZsJsPPZsGzwwsLwLmpwMDw",
	}
	got := determineGroupBadgeTotalPriority(rucksacks)
	want := 70

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}

