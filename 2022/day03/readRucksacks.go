package main

import (
	"bufio"
	"log"
	"os"
)

// ToDo: Make this a lib func like readFileLinesToStrings
func readRucksacks(filename string) []string {
	rucksacks := make([]string, 0)

	readFile, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		rucksack := fileScanner.Text()

		rucksacks = append(rucksacks, rucksack)
	}

	return rucksacks
}
