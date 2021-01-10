package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func partA(input []string) {

	stringsWithCorrectCharCount := 0

	for _, item := range input {
		// Split item into 3 parts
		temp := strings.Split(item, " ")

		// work out upper and lower limits in intger format
		limits := strings.Split(temp[0], "-")
		lowerLim, _ := strconv.Atoi(limits[0])
		upperLim, _ := strconv.Atoi(limits[1])

		charToFind := strings.Trim(temp[1], ":")
		stringToSearch := temp[2]

		found := strings.Count(stringToSearch, charToFind)

		if found >= lowerLim && found <= upperLim {
			stringsWithCorrectCharCount++
		}
	}

	println("String Count: ", stringsWithCorrectCharCount)
}

func partB(input []string) {

	validstring := 0

	for _, item := range input {
		// Split item into 3 parts
		temp := strings.Split(item, " ")
		positions := strings.Split(temp[0], "-")
		position1, _ := strconv.Atoi(positions[0])
		position2, _ := strconv.Atoi(positions[1])

		temp2 := strings.Trim(temp[1], ":")
		charToFind := temp2[0]
		stringToSearch := temp[2]

		eq1 := stringToSearch[position1-1] == charToFind
		eq2 := stringToSearch[position2-1] == charToFind

		if (eq1 || eq2) && (eq1 != eq2) {
			validstring++
		}
	}

	println("Part 2 String Count: ", validstring)
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
