package main

import (
	"fmt" // library for printing stuff to the console
    "strconv" // library for converting strings into other datatypes
    "strings" // library for string manipulation
    "adventofcode/util"
)

func main(){
	fmt.Println("Starting day two programm")
	scanner, file := util.ReadInputfile("daytwo_input.txt")
	defer file.Close()
	var line string
	safeCounter := 0
	fmt.Println("Calculating the safe reports")
	for scanner.Scan() {
		line = scanner.Text()
		if isSafe(convertStringArrayToIntArray(strings.Split(line, " "))) {
			safeCounter++
		}
	}
	fmt.Printf("There are: %d safe reports\n", safeCounter)
}

func isSafe(slice []int) bool {
	fmt.Println(slice)
	ascending := slice[0] < slice[1]
	for i := 0; i < len(slice) - 1; i++ {
		if slice[i] < slice[i + 1] != ascending {
			return false
		}
		difference := util.Abs(slice[i] - slice[i + 1])
		if difference < 1 || difference > 3 {
			fmt.Println("false")
			return false
		}
	}
	fmt.Println("true")
	return true
}

func convertStringArrayToIntArray(stringSlice []string) []int {
	var intSlice []int
	for i := 0; i < len(stringSlice); i++ {
		stringValueAsInt, err := strconv.Atoi(stringSlice[i])
		if err != nil {
			panic(err)
		}
		intSlice = append(intSlice, stringValueAsInt)
	}
	return intSlice
}