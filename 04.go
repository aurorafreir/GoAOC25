package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"example.com/utils"
)

func countCharacters(inputStr string, charToFind string) int {
	charCount := 0
	for _, i := range inputStr {
		if string(i) == charToFind {
			charCount++
		}
	}
	return charCount
}

func flattenSlice(inputSlice []string) (outputString string) {
	//
	outputString = strings.Join(inputSlice, "")

	return outputString
}

func boxFilter(inputSlice []string, xRadius int, yRadius int, x int, y int, flatten bool) []string {
	//
	var outputSlice []string
	yMin, yMax := max(x-xRadius, 0), min(x+xRadius+1, len(inputSlice))
	for _, Y := range inputSlice[yMin:yMax] {
		xMin, xMax := max(y-yRadius, 0), min(y+yRadius+1, len(Y))
		// fmt.Println(Y[xMin:xMax])
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

	for indexY, itemY := range data {
		// fmt.Println(indexY, item)
		for indexX, itemX := range itemY {
			if string(itemX) == "@" {
				outputString := flattenSlice(boxFilter(data, 1, 1, indexX, indexY, true))
				fmt.Println(indexY, indexX, outputString)
				if (countCharacters(outputString, "@") - 1) < 4 {
					fmt.Println(countCharacters(outputString, "@"))
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
