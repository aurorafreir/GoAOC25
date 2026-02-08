package main

import (
	"fmt"
	"strconv"
	"strings"

	"example.com/utils"
)

// Rotates a slice 90 degrees clockwise
func rotateStringSlice90(inputSlice [][]string) (output [][]string, err error) {
	// Rotates a slice
	yLen := len(inputSlice)
	output = make([][]string, len(inputSlice[0])+1)
	for y := range inputSlice { // Initialize slice of slices
		output[y] = make([]string, yLen)
	}
	for y := range inputSlice {
		for x := range yLen {
			output[x][y] = inputSlice[y][x]
		}
	}
	return output, err
}

func d6p1() (out int, err error) {
	testing := true
	data, err := utils.AOCFileReadToSlice(false, 6)
	utils.Check(err)

	sliceOfItems := [][]string{}
	for _, line := range data {
		sliceOfLine, _ := utils.SplitStrIntoArbitraryLength(line, 4)
		sliceOfItems = append(sliceOfItems, sliceOfLine)
	}

	if testing {
		fmt.Println("sliceOfItems:", sliceOfItems)
	}
	rotatedSliceOfItems, _ := rotateStringSlice90(sliceOfItems)
	if testing {
		fmt.Println("rotatedSliceOfItems:", rotatedSliceOfItems)
	}

	total := 0
	for _, set := range rotatedSliceOfItems {
		numbers := set[0 : len(set)-2]
		sign := strings.ReplaceAll(set[len(set)-2], " ", "") // sign, removed whitespace
		setOutput, _ := strconv.Atoi(strings.ReplaceAll(numbers[0], " ", ""))
		for _, number := range numbers[1:] {
			numberInt, _ := strconv.Atoi(strings.ReplaceAll(number, " ", ""))
			switch sign {
			case "*":
				setOutput = setOutput * numberInt
			case "+":
				setOutput = setOutput + numberInt
			}
		}
		total += setOutput
	}

	return total, nil
}

func main() {
	fmt.Println(d6p1())
}
