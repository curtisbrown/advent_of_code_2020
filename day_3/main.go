package main

import (
	"bufio"
	"log"
	"os"
)

var lineLength int = 31

func partA(input []string) {

	treeCount := 0
	lineIdx := 0

	for i := 0; i < len(input)-1; i++ {

		lineIdx = (lineIdx + 3) % 31
		println("Line Index: ", lineIdx, "    Value: ", string(input[i+1][lineIdx]))
		if string(input[i+1][lineIdx]) == "#" {
			treeCount++
		}

	}
	println("String Count: ", treeCount)
}

func partB(input []string) {

	println("Part 2 String Count: ")
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	lines := []string{""}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	// Remove first entry as its empty
	lines = lines[1:]

	partA(lines)

	partB(lines)
}
