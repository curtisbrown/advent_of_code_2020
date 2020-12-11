package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func partA(nums []int) {

	found := false
	product := 0
	for idx, num := range nums {
		for idx2, num2 := range nums {
			if idx != idx2 {
				if num+num2 == 2020 {
					println("FOUND COMBO: Lines ", idx, " and ", idx2, "\nNumbers are :", num, " and ", num2)
					found = true
					product = num * num2
					break
				}
			}
		}
		if found {
			break
		}
	}

	println("Product : ", product)
}

func partB(nums []int) {

	found := false
	product := 0
	nums = nums[1:]
	for idx, num := range nums {
		for idx2, num2 := range nums {
			numToFind := 2020 - num - num2
			for idx3, num3 := range nums {
				if idx != idx2 && idx != idx3 {
					if num3 == numToFind {
						println("FOUND COMBO: Lines ", idx, " and ", idx2, " and ", idx3, "\nNumbers are :", num, " and ", num2, " and ", num3)
						found = true
						product = num * num2 * num3
						break
					}
				}
			}
			if found {
				break
			}
		}
		if found {
			break
		}
	}

	println("Product : ", product)
}

func main() {

	f, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	lines := []int{0}
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		lines = append(lines, num)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	partA(lines)

	partB(lines)

}
