package main

import (
	"bufio"
	"log"
	"os"
	"regexp"
	"strings"
)

func checkPassportType(entry string) bool {
	//byr (Birth Year)
	//iyr (Issue Year)
	//eyr (Expiration Year)
	//hgt (Height)
	//hcl (Hair Color)
	//ecl (Eye Color)
	//pid (Passport ID)
	//cid (Country ID)

	if strings.Count(entry, ":") == 8 || (strings.Count(entry, ":") == 7 && !strings.Contains(entry, "cid")) {
		return true
	}

	return false
}

func checkPassportStrict(entry string, f1 *os.File) bool {

	regexMap := make(map[string]string)

	// byr (Birth Year) - four digits; at least 1920 and at most 2002.
	regexMap["byr"] = "^(19[2-9][0-9]|200[0-2])$"
	// iyr (Issue Year) - four digits; at least 2010 and at most 2020.
	regexMap["iyr"] = "^(201[0-9]|2020)$"
	// eyr (Expiration Year) - four digits; at least 2020 and at most 2030.
	regexMap["eyr"] = "^(202[0-9]|2030)$"
	// hgt (Height) - a number followed by either cm or in:
	// If cm, the number must be at least 150 and at most 193.
	// If in, the number must be at least 59 and at most 76.
	regexMap["hgt"] = "^(1[5-8][0-9]cm|19[0-3]cm)|(59in|6[0-9]in|7[0-6]in)$"
	// hcl (Hair Color) - a # followed by exactly six characters 0-9 or a-f.
	regexMap["hcl"] = "^[#][0-9a-f]{6}$"
	// ecl (Eye Color) - exactly one of: amb blu brn gry grn hzl oth.
	regexMap["ecl"] = "^amb|blu|brn|gry|grn|hzl|oth$"
	// pid (Passport ID) - a nine-digit number, including leading zeroes.
	regexMap["pid"] = "^[0-9]{9}$"

	// Remove leading whitespace
	entry = strings.TrimSpace(entry)

	// Split into fields
	dataPairs := strings.Split(entry, " ")

	valueMap := make(map[string]string)
	for _, item := range dataPairs {
		pair := strings.Split(item, ":")
		valueMap[pair[0]] = pair[1]
		matched, _ := regexp.MatchString(regexMap[pair[0]], pair[1])
		if !matched {
			return false
		}
	}

	_, err2 := f1.WriteString("byr:" + valueMap["byr"] + " iyr:" + valueMap["iyr"] + " eyr:" + valueMap["eyr"] + " hgt:" + valueMap["hgt"] + " hcl:" + valueMap["hcl"] + " ecl:" + valueMap["ecl"] + " pid:" + valueMap["pid"] + "\n")
	if err2 != nil {
		log.Fatal(err2)
	}

	return true
}

func sortArrayEntries(input []string) []string {

	var newTempArr []string
	var tempStr string

	emptyline := 0
	// Sort out the input into groups of complete data
	for _, item := range input {

		if item != "" {
			tempStr = tempStr + " " + item
		} else {
			emptyline++
			newTempArr = append(newTempArr, tempStr)
			tempStr = ""
		}
	}

	// Get last line of data after index limit has been reached
	if tempStr != "" {
		newTempArr = append(newTempArr, tempStr)
	}

	return newTempArr
}

func partA(input []string) {

	sortedArr := sortArrayEntries(input)
	count := 0
	// Process new set of entries
	for _, item := range sortedArr {
		result := checkPassportType(item)
		if result {
			count++
		}
	}

	println("valid count : ", count)

}

func partB(input []string) {

	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()

	f1, err := os.Create("outputSorted.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f1.Close()

	sortedArr := sortArrayEntries(input)
	count := 0
	// Process new set of entries
	for _, item := range sortedArr {
		println(item)
		result := checkPassportType(item)
		if result {
			if checkPassportStrict(item, f1) {
				_, err2 := f.WriteString(item + "\n")
				if err2 != nil {
					log.Fatal(err2)
				}
				count++
			}
		}
	}

	println("strict valid count : ", count)
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
