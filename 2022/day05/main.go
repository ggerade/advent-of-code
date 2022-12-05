package main

import (
	"fmt"
)

func determineTopsOfStacks9000(stacks []Stack, instructions []Instruction) string {
	// Follow each instruction
	for _, instruction := range instructions {
		for i := 0; i < instruction.move; i++ {
			from := stacks[instruction.from-1]
			popped := from[len(from)-1]
			stacks[instruction.from-1] = from[:len(from)-1]
			stacks[instruction.to-1] = append(stacks[instruction.to-1], popped)
		}
	}

	// Grab top of each stack to string
	tops := ""
	for _, stack := range stacks {
		tops += stack[len(stack)-1]
	}
	return tops
}

func determineTopsOfStacks9001(stacks []Stack, instructions []Instruction) string {
	// Move em'
	for _, instruction := range instructions {
		popped := make(Stack, 0)

		for i := 0; i < instruction.move; i++ {
			from := stacks[instruction.from-1]
			// pop and store in reverse
			popped = append(Stack{from[len(from)-1]}, popped...)
			stacks[instruction.from-1] = from[:len(from)-1]
		}
		//fmt.Printf("popped: %v\n", popped);
		// push what we popped
		stacks[instruction.to-1] = append(stacks[instruction.to-1], popped...)
	}

	// Grab top of each stack to string
	tops := ""
	for _, stack := range stacks {
		tops += stack[len(stack)-1]
	}
	return tops
}

func main() {
	stacks := readStacks("inputs/exampleStacks.txt")
	moves := readInstructions("inputs/exampleInstructions.txt")
	fmt.Printf("Example stacks with 9000: %s\n", determineTopsOfStacks9000(stacks, moves))
	stacks = readStacks("inputs/exampleStacks.txt")	
	fmt.Printf("Example stacks with 9001: %s\n", determineTopsOfStacks9001(stacks, moves))

	stacks2 := readStacks("inputs/challengeStacks.txt")
	moves2 := readInstructions("inputs/challengeInstructions.txt")
	fmt.Printf("Example stacks with 9000: %s\n", determineTopsOfStacks9000(stacks2, moves2))
	stacks2 = readStacks("inputs/challengeStacks.txt")
	fmt.Printf("Example stacks with 9001: %s\n", determineTopsOfStacks9001(stacks2, moves2))
}
