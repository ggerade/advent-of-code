package main

import (
	"fmt"
	"io"
	"log"
	"os"

	"github.com/icza/backscanner"
)

type Stack []string

func readStacks(filename string) []Stack {
	stacks := make([]Stack, 0)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := backscanner.New(file, int(fileInfo.Size()))
	skippedFirstLine := false
	for {
		line, _, err := scanner.LineBytes()
		if err != nil {
			if err != io.EOF {
				log.Fatal("Error:", err)
			}
			break
		}
		if !skippedFirstLine {
			skippedFirstLine = true
			continue
		}

		for i, pos := 0, 1; pos < len(line); i, pos = i+1, pos+4 {
			if line[pos] != ' ' {
				if len(stacks) < i+1 {
					stacks = append(stacks, make(Stack, 0))
				}
				stacks[i] = append(stacks[i], string(line[pos]))
			}
		}
	}

	return stacks
}
