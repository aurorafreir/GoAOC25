package main

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"

	"example.com/utils"
)

func lenItoa(i int) int {
	// From stackoverflow https://stackoverflow.com/a/68122831
	return len(strconv.Itoa(i))
}

func d2p1() int {
	// Basic file importing
	ex, err := os.Getwd()
	utils.Check(err)

	path := filepath.Join(ex, "testinput.txt")
	data, err := utils.ReadLines(path)
	utils.Check(err)

	// Get just the first line as a string
	singleString := data[0]

	// Split the input string by "," into each inputRange
	inputRanges := strings.Split(singleString, ",")
	for _, inputRange := range inputRanges {
		// Get individual Start and End ranges as strings
		startRange := strings.Split(inputRange, "-")[0]
		endRange := strings.Split(inputRange, "-")[1]

		// Convert Start and End strings into ints
		startRangeInt, err := strconv.Atoi(startRange)
		utils.Check(err)
		endRangeInt, err := strconv.Atoi(endRange)
		utils.Check(err)

		// Get start and end of ranges and loop through them
		fmt.Println("start:", startRangeInt, "end:", endRangeInt)
		// lenOfEndInt := lenItoa(endRangeInt)
		// fmt.Println(lenOfEndInt)
		// for i := startRangeInt; i <= endRangeInt; i++ {

		for i := startRangeInt; i <= endRangeInt; i++ {
			// Find looping numbers in each int
			// fmt.Println(i)
			lenOfInt := lenItoa(i)
			// fmt.Println(lenOfInt)
			var sliceOfStrs []string

			// Splits each integer into a Slice of ints of each
			for x := 0; x <= lenOfInt; x++ {
				for y := x; y < lenOfInt; y++ {
					if y == x {
						continue
					}
					slicedIntAsStr := strconv.Itoa(i)
					slicedStr := slicedIntAsStr[x:y]
					utils.Check(err)
					// No point in saving the single length items
					if len(slicedStr) > 1 {
						sliceOfStrs = append(sliceOfStrs, slicedStr)
					}
				}
			}
			// Sorts, and then removes duplicates from $sliceOfStrs, outputs to $sortedSliceOfStrs
			slices.Sort(sliceOfStrs)
			sortedSliceOfStrs := slices.Compact(sliceOfStrs)
			fmt.Println(sortedSliceOfStrs)
			// regexComp := strings.Join(sortedSliceOfStrs, "|")
			// regexComp := regexp.MustCompile(strings.Join(sortedSliceOfStrs, "|"))
			// fmt.Println(regexComp)
			// fmt.Println(len(regexComp.FindAllString(strconv.Itoa(i), -1)))
		}
	}
	return 0
}

func d2p2() int {
	return 0
}

// func main() {
// 	fmt.Println("Result A:", partA())
// 	fmt.Println("Result B:", partB())
// }
