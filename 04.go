package main

import (
	"os"
	"path/filepath"
	"strings"

	"example.com/utils"
)

func countCharacters(inputStr string, charToFind string) int {
	// Counts the number of matching characters in a string, should probably switch to a regex function
	charCount := 0
	for _, i := range inputStr {
		if string(i) == charToFind {
			charCount++
		}
	}
	return charCount
}

func flattenSlice(inputSlice []string) (outputString string) {
	// Flattens a slice into a single string
	outputString = strings.Join(inputSlice, "")

	return outputString
}

func boxFilter(inputSlice []string, xRadius int, yRadius int, x int, y int) []string {
	// Safe box filter given an input slice of Y (slice, vertical) and X (string, horizontal)
	// $xRadius for string search radius, $yRadius for slice search radius
	// Returns a slice with the box filter's search radius
	var outputSlice []string
	yMin, yMax := max(y-yRadius, 0), min(y+yRadius+1, len(inputSlice))
	for _, Y := range inputSlice[yMin:yMax] {
		xMin, xMax := max(x-xRadius, 0), min(x+xRadius+1, len(Y))
		outputSlice = append(outputSlice, Y[xMin:xMax])
	}

	return outputSlice
}

func d4p1() (int, error) {
	// Basic file importing
	ex, err := os.Getwd()
	utils.Check(err)

	path := filepath.Join(ex, "input", "04.txt")
	data, err := utils.ReadLines(path)
	utils.Check(err)

	maxAccessibleRolls := 0

	// loop through the Y and X ranges, check if each item is an @,
	// then run a box filter on each item, and check the @ count in surrounding items
	for indexY, itemY := range data {
		for indexX, itemX := range itemY {
			if string(itemX) == "@" {
				outputString := flattenSlice(boxFilter(data, 1, 1, indexX, indexY))
				if (countCharacters(outputString, "@") - 1) < 4 {
					maxAccessibleRolls++
				}
			}
		}
	}

	return maxAccessibleRolls, nil
}

func d4p2() (int, error) {
	return 0, nil
}
