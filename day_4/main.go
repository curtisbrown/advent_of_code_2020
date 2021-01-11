package main

import (
	"bufio"
	"log"
	"os"
)

func partA(input []string) {

}

func partB(input []string) {

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
