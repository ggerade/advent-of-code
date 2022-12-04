package main

import (
	"fmt"
)

type Assignment struct {
	begin int
	end   int
}

func determineFullyContainedPairs(assignments []Assignment) int {
	total := 0
	for i := 0; i < len(assignments); i += 2 {
		first := assignments[i]
		second := assignments[i+1]
		if first.begin <= second.begin && first.end >= second.end ||
			second.begin <= first.begin && second.end >= first.end {
			total++
		}
	}
	return total
}

func determineOverlappedPairs(assignments []Assignment) int {
	total := 0
	for i := 0; i < len(assignments); i += 2 {
		first := assignments[i]
		second := assignments[i+1]
		if second.begin >= first.begin && second.begin <= first.end ||
		  second.end >= first.begin && second.begin <= first.end {
			total++
		}
	}
	return total
}

func main() {
	exampleAssignments := readAssignments("inputs/example.txt")
	fmt.Printf("Example #1 Total contained : %d\n", determineFullyContainedPairs(exampleAssignments))
	fmt.Printf("Example #2 Total overlapped: %d\n", determineOverlappedPairs(exampleAssignments))

	challengeAssignments := readAssignments("inputs/challenge.txt")
	fmt.Printf("Challenge #1 Total contained : %d\n", determineFullyContainedPairs(challengeAssignments))
	fmt.Printf("Challenge #2 Total overlapped: %d\n", determineOverlappedPairs(challengeAssignments))
}
