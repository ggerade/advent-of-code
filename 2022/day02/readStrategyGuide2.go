package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func convertKeyToPlay2(key string) Play {
	switch key {
	case "A": return ROCK
	case "B": return PAPER
	}
	return SCISSORS
}

func convertKeyToGoal(key string) Goal {
	switch key {
	case "X": return LOSS
	case "Y": return DRAW
	}
	return WIN
}

func readStrategyGuide2(filename string) []Strategy2 {
	strategyGuide := make([]Strategy2, 0)

	readFile, err := os.Open(filename)
	if err != nil {
			log.Fatal(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		keys := strings.Split(line, " ")
		
		strategy := Strategy2 {
			opponent: convertKeyToPlay2(keys[0]),
			goal: convertKeyToGoal(keys[1]),
		}
		strategyGuide = append(strategyGuide, strategy)
	}

	// Better to defer this I think?
	readFile.Close()

	return strategyGuide;
}
