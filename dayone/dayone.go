package main

import (
    "fmt" // library for printing stuff to the console
    "log" 
    "os" // library for executing os calls
    "bufio" // library for reading file input
    "strconv" // library for converting strings into other datatypes
    "strings" // library for string manipulation
)

func main() {
    fmt.Println("Starting day one programm")
    // open file

    f, err := os.Open("dayone_input.txt")
    if err != nil {
        log.Fatal(err)
    }
    // remember to close the file at the end of the program
    defer f.Close()

    // read the file line by line using scanner
    scanner := bufio.NewScanner(f)
    var leftList []int
    var rightList []int
    var line []string
    for scanner.Scan() {
        // do something with a line
        line = strings.Split(scanner.Text(), "   ") //split into left and right list
        leftListValue, err := strconv.Atoi(line[0]) // converting the string value into an integer
        if err != nil {
            panic(err) // panic on any error
        }
        leftList = append(leftList, leftListValue) // append the parsed value to the left list
        rightListValue, err := strconv.Atoi(line[1]) // converting the string value into an integer
        if err != nil {
            panic(err) // panic on any error
        }
        rightList = append(rightList, rightListValue) // append the parsed value to the right list
    }

    if err := scanner.Err(); err != nil {
        log.Fatal(err)
    }

    fmt.Println("Sorting the lists") // sorting lists for pairing the lowest item in the left with the lowest in the right (...) on the same index 
    leftList = dayone_bubblesort(leftList)
    rightList = dayone_bubblesort(rightList)
    fmt.Println("Sorting done")

    partone(leftList, rightList)
    parttwo(leftList, rightList)
}

// solve part one
func partone (leftList []int, rightList []int){
    fmt.Println("Calculating total distance")
    listsLength := len(leftList)
    totalDistance := 0
    for i := 0; i < listsLength; i++ {
        totalDistance += partone_abs(leftList[i] - rightList[i])
    }
    fmt.Printf("The total distance is: %d\n", totalDistance)
}

// solve part two
func parttwo(leftList []int, rightList []int) {
    similarityScore := 0;
    listsLength := len(rightList)
    countMap := make(map[int]int) //initialise a map with a pointer (var m map[string]int would just behave like an empty map on read but would cause runtime panic when writing)
    fmt.Println("Calculating the similarity score")
    for i := 0; i < listsLength; i++ {
        countMap[rightList[i]]++
    }
    for i := 0; i < listsLength; i++ {
        similarityScore += leftList[i] * countMap[leftList[i]]
    }
    fmt.Printf("The total similarity score is: %d\n", similarityScore)

}

// implementing bubblesort just for practice reasons
func dayone_bubblesort(slice []int) []int {
    var temp int // declare temporary swap variable
    for rightIndex := len(slice); 1 < rightIndex; rightIndex-- {
        for leftIndex := 0; leftIndex < rightIndex - 1; leftIndex++ { 
            if slice[leftIndex] > slice[leftIndex + 1] {
                temp = slice[leftIndex]
                slice[leftIndex] = slice[leftIndex + 1]
                slice[leftIndex + 1] = temp
            }
        }
    }
    return slice
}

func partone_abs(number int) int {
    if number < 0 {
        return number * -1
    }
    return number
}