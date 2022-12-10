package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

const (
	DIR = iota
	FILE
)

type Node struct {
	name     string
	tipe     int
	size     int
	parent   *Node
	children []*Node
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func readOutput(filename string) Node {
	root := Node{
		name:     "/",
		tipe:     DIR,
		size:     0,
		children: make([]*Node, 0),
	}

	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	cwd := &root
	for fileScanner.Scan() {
		line := fileScanner.Text()
		//fmt.Println(line)

		if strings.HasPrefix(line, "$ cd /") {
			fmt.Printf("cd to /\n")
			cwd = &root
		} else if strings.HasPrefix(line, "$ cd ..") {
			fmt.Printf("cd .. to %s\n", cwd.parent.name)
			cwd = cwd.parent
		} else if strings.HasPrefix(line, "$ cd ") {
			dir := line[5:]
			for _, child := range cwd.children {
				if child.tipe == DIR && dir == child.name {
					fmt.Printf("cd to %s\n", child.name)
					cwd = child
					break
				}
			}
		} else if line == "$ ls" {
		} else {
			regex := regexp.MustCompile(`(.*) (.*)`)
			subs := regex.FindStringSubmatch(line)
			if subs[1] == "dir" {
				cwd.children = append(cwd.children, &Node{
					name:   subs[2],
					tipe:   DIR,
					parent: cwd,
					size:   0,
				})
			} else {
				cwd.children = append(cwd.children, &Node{
					name:   subs[2],
					tipe:   FILE,
					parent: cwd,
					size:   atoi(subs[1]),
				})
			}
		}
	}
	return root
}

func calcDirSize(node *Node) int {
	for _, child := range node.children {
		if child.tipe == DIR {
			node.size += calcDirSize(child)
		} else {
			node.size += child.size
		}
	}
	return node.size
}

func findBigDirs(node Node, total *int) int {
	for _, child := range node.children {
		if child.tipe == DIR {
			if child.size <= 100000 {
				*total += child.size
			}
			node.size += findBigDirs(*child, total)
		}
		//  else {
		// 	if child.size >= 100000 {
		// 		*total += child.size
		// 	}
		// }
	}
	return *total
}

func findSmallestDirBigger(node Node, tofree int, smallest *int) int {
	for _, child := range node.children {
		if child.tipe == DIR {
			if child.size > tofree && child.size < *smallest {
				*smallest = child.size
			}
			findSmallestDirBigger(*child, tofree, smallest)
		}
	}
	return *smallest
}

func main() {
	// root := readOutput("inputs/example.txt")
	// size := calcDirSize(&root)
	// fmt.Printf("Root dir size %d\n", size)
	// total := 0
	// fmt.Printf("Example Big Dirs: %d\n", findBigDirs(root, &total))

	root := readOutput("inputs/challenge.txt")
	size := calcDirSize(&root)
	fmt.Printf("Root dir size %d\n", size)
	total := 0
	fmt.Printf("Challenge 1 Big Dirs: %d\n", findBigDirs(root, &total))

	remain := 70000000 - size
	tofree := 30000000 - remain
	smallest := math.MaxInt
	fmt.Printf("Challenge 2 Smallest Dirs: %d\n", findSmallestDirBigger(root, tofree, &smallest))
}
