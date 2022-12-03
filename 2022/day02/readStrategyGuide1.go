package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func convertKeyToPlay(key string) Play {
	switch key {
	case "A": return ROCK
	case "B": return PAPER
	case "X": return ROCK
	case "Y": return PAPER
	}
	return SCISSORS
}

func readStrategyGuide1(filename string) []Strategy1 {
	strategyGuide := make([]Strategy1, 0)

	readFile, err := os.Open(filename)
	if err != nil {
			fmt.Println(err)
			// and???
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
		keys := strings.Split(line, " ")
		
		strategy := Strategy1 {
			opponent: convertKeyToPlay(keys[0]),
			you: convertKeyToPlay(keys[1]),
		}
		strategyGuide = append(strategyGuide, strategy)
	}

	// Better to defer this I think?
	readFile.Close()

	return strategyGuide;
}
