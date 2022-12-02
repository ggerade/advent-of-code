package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode/utf8"
)

func fileTo2DimensionalSliceInts(filename string) [][]int {
	outer := make([][]int, 0)
	inner := make([]int, 0)

	readFile, err := os.Open(filename)
	if err != nil {
			fmt.Println(err)
			// and???
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text();
		if (utf8.RuneCountInString(line) > 0) {
			n, err := strconv.Atoi(line)
			if err != nil {
				fmt.Println(err)
			}	else {
				inner = append(inner, n)
			}
		} else {
			outer = append(outer, inner)
			inner = make([]int, 0)			
		}
	}
	outer = append(outer, inner)

	// Better to defer this I think?
	readFile.Close()

	return outer;
}
