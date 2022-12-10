package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Instruction struct {
	move int
	from int
	to   int
}

func readInstructions(filename string) []Instruction {
	instructions := make([]Instruction, 0)

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
    // move 1 from 2 to 1
		regex := regexp.MustCompile(`move (\d.*) from (\d.*) to (\d.*)`)
    subs := regex.FindStringSubmatch(line)
	
		instructions = append(instructions, Instruction{move: atoi(subs[1]), from: atoi(subs[2]), to: atoi(subs[3])})
	}

	return instructions
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return i;
}