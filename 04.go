package main

import (
	"fmt"
	"reflect"

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

func d4p1() (int, error) {
	data, err := utils.AOCFileReadToSlice(false, 4)
	utils.Check(err)

	maxAccessibleRolls := 0

	// loop through the Y and X ranges, check if each item is an @,
	// then run a box filter on each item, and check the @ count in surrounding items
	for indexY, itemY := range data {
		for indexX, itemX := range itemY {
			if string(itemX) == "@" {
				outputString := utils.FlattenSlice(utils.BoxFilter(data, 1, 1, indexX, indexY))
				if (countCharacters(outputString, "@") - 1) < 4 {
					maxAccessibleRolls++
				}
			}
		}
	}

	return maxAccessibleRolls, nil
}

func boxFilterEachItemAndReturnNewSlice(currentSlice []string) (outputSlice []string, modificationsMade int, err error) {
	// loop through the Y and X ranges, check if each item is an @,
	// then run a box filter on each item, and check the @ count in surrounding items
	// Then replaces the @ with a . if it has less than 4 @s in surrounding box filter
	for indexY, itemY := range currentSlice {
		newString := ""
		for indexX, itemX := range itemY {
			if string(itemX) == "@" {
				outputString := utils.FlattenSlice(utils.BoxFilter(currentSlice, 1, 1, indexX, indexY))
				if (countCharacters(outputString, "@") - 1) < 4 {
					modificationsMade++
					newString = newString + "."
				} else {
					newString = newString + "@"
				}
			} else {
				newString = newString + "."
			}
		}
		outputSlice = append(outputSlice, newString)
	}
	return outputSlice, modificationsMade, nil
}

func d4p2() (int, error) {
	data, err := utils.AOCFileReadToSlice(false, 4)
	utils.Check(err)

	maxAccessibleRolls := 0
	previousSlice := data

	for range 100 { // Max recursion depth
		newSlice, modificationsMade, _ := boxFilterEachItemAndReturnNewSlice(previousSlice)
		maxAccessibleRolls = maxAccessibleRolls + modificationsMade
		if reflect.DeepEqual(previousSlice, newSlice) { // Clear loop and return if previousSlice and newSlice are the same
			return maxAccessibleRolls, nil
		} else {
			previousSlice = newSlice
		}
	}
	fmt.Println("hit set max recursion depth, up recursion depth :)")

	return maxAccessibleRolls, nil
}
