package main

import (
	"adventofcode/util"
	"bufio"
	"fmt"
	"regexp"  // library for regular expression functions
	"strconv" // library for converting strings into other datatypes
	"strings" // library for string manipulation
)

func main() {
	fmt.Println("Staring day three program")
	scanner, file := util.ReadInputfile("daythree_input.txt")
	defer file.Close()
	solveDayThree(scanner)
}

func solveDayThree(scanner *bufio.Scanner) {
	fmt.Println("Calculating matches for partone")
	var matches []string
	var partTwoMatches []string
	instructionsEnabled := true
	partOneRegexMultiplicationPattern := regexp.MustCompile(`mul\(\d+,\d+\)`)
	partTwoRegexMultiplicationPattern := regexp.MustCompile(`don't\(\)|do\(\)|mul\(\d+,\d+\)`)
	for scanner.Scan() {
		// Get matches for part one
		partOneLineMatches := partOneRegexMultiplicationPattern.FindAllString(scanner.Text(), -1)
		matches = append(matches, partOneLineMatches...)
		// Get matches for part two
		partTwoLineMatches := partTwoRegexMultiplicationPattern.FindAllString(scanner.Text(), -1)
		for _, a := range partTwoLineMatches {
			if a == "don't()" {
				instructionsEnabled = false
				continue
			} else if a == "do()" {
				instructionsEnabled = true
				continue
			}
			if instructionsEnabled {
				partTwoMatches = append(partTwoMatches, a)
			}
		}

	}
	// Calculate sum for part one
	calculateSum(matches)
	// Calculate sum for part two
	calculateSum(partTwoMatches)
}

func calculateSum(matches []string) {
	sumOfResults := 0
	regexDigitsPattern := regexp.MustCompile(`(\d+,\d+)`)
	for _, a := range matches {
		multDigits := strings.Split(regexDigitsPattern.FindAllString(a, -1)[0], ",")
		firstDigit, err := strconv.Atoi(multDigits[0])
		if err != nil {
			panic(err)
		}
		secondDigit, err := strconv.Atoi(multDigits[1])
		if err != nil {
			panic(err)
		}
		sumOfResults += firstDigit * secondDigit
	}
	fmt.Printf("The sum of the multiplications is %d\n", sumOfResults)
}
