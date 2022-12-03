package main

import (
	"fmt"
	"strings"
)

func splitRucksackCompartments(rucksack string) (string, string) {
	return rucksack[0:(len(rucksack) / 2)], rucksack[len(rucksack)/2:]
}

// Hmm, maybe using a rune here is more trouble than its worth?
func findFirstDuplicate(compartment1 string, compartment2 string) rune {
	for _, r := range compartment1 {
		if strings.Contains(compartment2, string(r)) {
			return r
		}
	}
	// kind of lame, but one duplicate per expected in first challenge
	return '0'
}

func determinePriority(gift rune) int {
	if gift <= 'Z' {
		return int(gift) - 38
	} else {
		return int(gift) - 96
	}
}

func determineTotalPriority(rucksacks []string) int {
	total := 0
	for _, rucksack := range rucksacks {
		// split compartments
		compartment1, compartment2 := splitRucksackCompartments(rucksack)
		// find elf's error
		dupe := findFirstDuplicate(compartment1, compartment2)
		// determine priority
		priority := determinePriority(dupe)

		total += priority
	}
	return total
}

func findGroupBadge(rucksack1 string, rucksack2 string, rucksack3 string) rune {
	for _, r := range rucksack1 {
		if strings.Contains(rucksack2, string(r)) && strings.Contains(rucksack3, string(r)) {
			return r
		}
	}
	// kind of lame, but one badge across all expected in first challenge
	return '0'
}

func determineGroupBadgeTotalPriority(rucksacks []string) int {
	total := 0
	for i := 0; i < len(rucksacks); i += 3 {
		// find group's badge
		badge := findGroupBadge(rucksacks[i], rucksacks[i+1], rucksacks[i+2])
		total += determinePriority(badge)
	}
	return total
}

func main() {
	rucksacks := readRucksacks("inputs/challenge.txt")

	fmt.Printf("Challenge #1 Total priority: %d\n", determineTotalPriority(rucksacks))

	fmt.Printf("Challenge #2 Total group badge priority: %d\n", determineGroupBadgeTotalPriority(rucksacks))
}
