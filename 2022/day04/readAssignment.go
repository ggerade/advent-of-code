package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
)

// ToDo: Make this a lib func like readFileLinesToStrings - maybe make it more functional and generic
func readAssignments(filename string) []Assignment {
	assignments := make([]Assignment, 0)

	readFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		// and???
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()
    regex := regexp.MustCompile(`(\d.*)-(\d.*),(\d.*)-(\d.*)`)
    subs := regex.FindStringSubmatch(line)
		assignments = appendAssignment(assignments, subs[1], subs[2])
		assignments = appendAssignment(assignments, subs[3], subs[4])
	}

	return assignments
}

func appendAssignment(assignments []Assignment, begin string, end string) []Assignment {
	l, errl := strconv.Atoi(begin)
	if errl != nil {
		fmt.Println(errl)
	}
	r, errr := strconv.Atoi(end)
	if errr != nil {
		fmt.Println(errr)
	}

	assignments = append(assignments, Assignment{begin: l, end: r})
	return assignments
}
