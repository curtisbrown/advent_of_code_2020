package main

import (
	"bufio"
	"log"
	"os"
)

var lineLength int = 31

func treeCount(input []string, xCoord int, yCoord int) int {

	treeCount := 0
	lineIdx := 0
	for i := 0; i < len(input)-yCoord; i = i + yCoord {

		lineIdx = (lineIdx + xCoord) % 31
		// println("(", lineIdx, ",", i, ")", "    Value: ", string(input[i+yCoord][lineIdx]))
		if string(input[i+yCoord][lineIdx]) == "#" {
			treeCount++
		}

	}
	println("String Count: ", treeCount)
	return treeCount
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

	ret1 := treeCount(lines, 1, 1)
	ret2 := treeCount(lines, 3, 1)
	ret3 := treeCount(lines, 5, 1)
	ret4 := treeCount(lines, 7, 1)
	ret5 := treeCount(lines, 1, 2)

	println("Total Trees: ", ret1*ret2*ret3*ret4*ret5)
}
