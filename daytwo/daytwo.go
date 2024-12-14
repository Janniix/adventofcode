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
	totalSafeCounter := 0
	fmt.Println("Calculating the safe reports")
	for scanner.Scan() {
		line = scanner.Text()
		lineAsIntegerArray := convertStringArrayToIntArray(strings.Split(line, " "))
		if isSafe(lineAsIntegerArray) {
			safeCounter++
			totalSafeCounter++
		} else if isSafeAfterRemoveingASingleBadLevel(lineAsIntegerArray) {
			totalSafeCounter++
		}
	}
	fmt.Printf("There are: %d safe reports without the dampener\n", safeCounter)
	fmt.Printf("There are %d total safe reports with the dampener\n", totalSafeCounter)
}

func isSafe(slice []int) bool {
	ascending := slice[0] < slice[1]
	for i := 0; i < len(slice) - 1; i++ {
		if slice[i] < slice[i + 1] != ascending {
			return false
		}
		difference := util.Abs(slice[i] - slice[i + 1])
		if difference < 1 || difference > 3 {
			return false
		}
	}
	return true
}

func isSafeAfterRemoveingASingleBadLevel(slice []int) bool {
	for i := 0; i < len(slice); i++ {
		if isSafe(removeElementAtIndex(slice, i)) {
			return true
		}
	}
	return false
}

func removeElementAtIndex(slice []int, index int) []int {
	// create new slice and copy values of original 
	// note: could also kinda be done with the the append() function, but this would also change the original slice in some cases (e.g. if slice is as big or bigger then the array created by the append function) as a side effect 
	newSlice := make([]int, len(slice)-1)
	copy(newSlice[:index], slice[:index])
	copy(newSlice[index:], slice[index+1:])
	return newSlice
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