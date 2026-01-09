package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"example.com/generics"
)

func partA() int {
	// Basic file importing
	ex, err := os.Getwd()
	generics.Check(err)

	path := filepath.Join(ex, "testinput.txt")
	data, err := generics.ReadLines(path)
	generics.Check(err)

	// Get just the first line as a string
	singleString := data[0]
	// fmt.Println(singleString)

	// Split the input string by "," into each inputRange
	inputRanges := strings.Split(singleString, ",")
	for _, inputRange := range inputRanges {
		// Get individual Start and End ranges as strings
		startRange := strings.Split(inputRange, "-")[0]
		endRange := strings.Split(inputRange, "-")[1]

		// Convert Start and End strings into ints
		startRangeInt, err := strconv.Atoi(startRange)
		generics.Check(err)
		endRangeInt, err := strconv.Atoi(endRange)
		generics.Check(err)

		fmt.Println("start:", startRangeInt, "end:", endRangeInt)
		for i := startRangeInt; i <= endRangeInt; i++ {
			fmt.Println(i)
		}
	}
	return 0
}

func partB() int {
	return 0
}

func main() {
	fmt.Println("Result A:", partA())
	fmt.Println("Result B:", partB())
}
