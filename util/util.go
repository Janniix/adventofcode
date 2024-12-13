package util

import (
    "os" // library for executing os calls
    "bufio" // library for reading file input
    "log" 
)

func ReadInputfile(filepath string) (*bufio.Scanner, *os.File) {
	// open file

    f, err := os.Open(filepath)
    if err != nil {
        log.Fatal(err)
    }
    scanner := bufio.NewScanner(f)
	return scanner, f
}