package main

import (
	"fmt"
	"strconv" // library for converting strings into other datatypes
    "strings" // library for string manipulation
	"regexp" // library for regular expression functions
    "adventofcode/util"
)

func main() {
	fmt.Println("Staring day three program")
	scanner, file := util.ReadInputfile("daythree_input.txt")
	defer file.Close()
	var matches []string
	regexMultiplicationPattern :=  regexp.MustCompile(`mul\(\d+,\d+\)`)
	for scanner.Scan() {
		lineMatches := regexMultiplicationPattern.FindAllString(scanner.Text(), -1)
		for _, a := range(lineMatches) {
			matches = append(matches,a)
		}
	}
	calculateSum(matches)
	
}

func calculateSum(matches []string) {
	sumOfResults := 0
	regexDigitsPattern :=  regexp.MustCompile(`(\d+,\d+)`)
	for _, a := range(matches) {
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