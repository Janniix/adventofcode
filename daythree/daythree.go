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
	//`don't\(\)|do\(\)|mul\(\d+,\d+\)`gm
	fmt.Println("Staring day three program")
	scanner, file := util.ReadInputfile("daythree_input.txt")
	defer file.Close()
	partone(scanner)
}

func partone(scanner *bufio.Scanner) {
	fmt.Println("Calculating matches for partone")
	var matches []string
	var partTwoMatches []string
	instructionsEnabled := true
	partOneRegexMultiplicationPattern := regexp.MustCompile(`mul\(\d+,\d+\)`)
	partTowRegexMultiplicationPattern := regexp.MustCompile(`don't\(\)|do\(\)|mul\(\d+,\d+\)`)
	for scanner.Scan() {
		partOneLineMatches := partOneRegexMultiplicationPattern.FindAllString(scanner.Text(), -1)
		for _, a := range partOneLineMatches {
			matches = append(matches, a)
		}
		partTwoLineMatches := partTowRegexMultiplicationPattern.FindAllString(scanner.Text(), -1)
		for _, a := range partTwoLineMatches {
			fmt.Println(a)
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
	calculateSum(matches)
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
