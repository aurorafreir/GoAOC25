package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"example.com/utils"
)

// Rotates a slice 90 degrees clockwise
func rotateStringSlice90(inputSlice [][]string) (output [][]string, err error) {
	// Rotates a slice
	yLen := len(inputSlice)
	output = make([][]string, len(inputSlice[0]))
	for y := range output { // Initialize slice of slices
		output[y] = make([]string, yLen+1)
	}
	for y := range yLen {
		for x := range len(inputSlice[0]) {
			output[x][y] = inputSlice[y][x]
		}
	}
	return output, err
}

func d6p1() (out int, err error) {
	testing := false
	data, err := utils.AOCFileReadToSlice(testing, 6)
	utils.Check(err)

	nonSpaceIndexes := make([]bool, len(data[0]))
	for _, line := range data {
		for index, char := range line {
			if !(string(char) == " ") {
				nonSpaceIndexes[index] = true
			}
		}
	}

	cleanSliceOfSlices := [][]string{}
	for _, line := range data {
		cleanSlice := []string{}
		prevIndex := 0
		for index := range line {
			if !nonSpaceIndexes[index] {
				cleanSlice = append(cleanSlice, line[prevIndex:index])
				prevIndex = index
			}
			if index == len(line)-1 {
				cleanSlice = append(cleanSlice, line[prevIndex:index+1])
			}
		}
		cleanSliceOfSlices = append(cleanSliceOfSlices, cleanSlice)
	}

	if testing {
		fmt.Println("sliceOfItems:", cleanSliceOfSlices)
	}
	rotatedSliceOfItems, _ := rotateStringSlice90(cleanSliceOfSlices)
	if testing {
		fmt.Println("rotatedSliceOfItems:", rotatedSliceOfItems)
	}

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
		out += setOutput
	}

	return out, nil
}

func d6p2() (out int, err error) {
	testing := false
	data, err := utils.AOCFileReadToSlice(testing, 6)
	utils.Check(err)

	nonSpaceIndexes := make([]bool, len(data[0]))
	for _, line := range data {
		for index, char := range line {
			if !(string(char) == " ") {
				nonSpaceIndexes[index] = true
			}
		}
	}

	cleanSliceOfSlices := [][]string{}
	for _, line := range data {
		cleanSlice := []string{}
		prevIndex := 0
		for index := range line {
			if !nonSpaceIndexes[index] {
				cleanSlice = append(cleanSlice, line[prevIndex:index])
				prevIndex = index
			}
			if index == len(line)-1 {
				cleanSlice = append(cleanSlice, line[prevIndex:index+1])
			}
		}
		cleanSliceOfSlices = append(cleanSliceOfSlices, cleanSlice)
	}

	if testing {
		fmt.Println("sliceOfItems:", cleanSliceOfSlices)
	}
	rotatedSliceOfItems, _ := rotateStringSlice90(cleanSliceOfSlices)
	slices.Reverse(rotatedSliceOfItems) // Input is read right to left after rotating
	if testing {
		fmt.Println("rotatedSliceOfItems:", rotatedSliceOfItems)
	}

	for _, set := range rotatedSliceOfItems {
		numbers := set[0 : len(set)-2]
		sign := strings.ReplaceAll(set[len(set)-2], " ", "") // sign, removed whitespace

		// Converts each column into it's own 2d slice, rotates it, and converts back to strings
		sliceSliceOfOutput := [][]string{}
		for _, number := range numbers {
			sliceSliceOfOutput = append(sliceSliceOfOutput, strings.Split(number, ""))
		}
		rotatedSubSliceOfItems, _ := rotateStringSlice90(sliceSliceOfOutput)
		slices.Reverse(rotatedSubSliceOfItems)

		setOutput, _ := strconv.Atoi(strings.ReplaceAll(strings.Join(rotatedSubSliceOfItems[0], ""), " ", "")) // Set starting output number to the first item
		for _, number := range rotatedSubSliceOfItems[1:] {
			numberInt, _ := strconv.Atoi(strings.ReplaceAll(strings.Join(number, ""), " ", "")) // Combines to string and converts to int
			if numberInt == 0 {                                                                 // Ignore empty slices (Should fix rotation code, but this works as a simple fix)
				continue
			}
			switch sign {
			case "*":
				setOutput = setOutput * numberInt
			case "+":
				setOutput = setOutput + numberInt
			}
		}
		out += setOutput
	}

	return out, nil
}
