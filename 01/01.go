package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"example.com/utils"
)

func directionAndIncrement(line string) (string, int, error) {
	direction := line[0:1]
	num := line[1:]
	intconv, err := strconv.Atoi(num)
	return direction, intconv, err
}

func partA() int {
	currentVal := 50
	zeroCounts := 0

	ex, err := os.Getwd()
	utils.Check(err)

	path := filepath.Join(ex, "input.txt")
	data, err := utils.ReadLines(path)
	utils.Check(err)
	for _, line := range data {
		direction, increment, err := directionAndIncrement(line)
		utils.Check(err)

		switch direction {
		case "L":
			currentVal -= increment
		case "R":
			currentVal += increment
		}

		if currentVal%100 == 0 {
			zeroCounts += 1
		}

	}
	return zeroCounts
}

func partB() int {
	previousVal := 50
	currentVal := 50
	zeroCounts := 0

	ex, err := os.Getwd()
	utils.Check(err)

	path := filepath.Join(ex, "input.txt")
	data, err := utils.ReadLines(path)
	utils.Check(err)
	for _, line := range data {

		// Split the line (e.g. "L68") into $direction ("L") and $strnum ("68"),
		// then converts $num to an integer
		direction, increment, err := directionAndIncrement(line)
		utils.Check(err)

		switch direction {
		case "L":
			currentVal -= increment
		case "R":
			currentVal += increment
		}

		// fmt.Println(direction, num, "prev:", previousVal, "current:", currentVal)

		// Avoids double hits when the previous answer was 0
		previousValTemp := previousVal
		if previousVal == 0 && currentVal < 0 {
			previousValTemp = -1
		} else if previousVal == 0 && currentVal > 0 {
			previousValTemp = 1
		}

		// Loops through all ints between previousVal and currentVal to see how many zero hits there are
		for i := min(currentVal, previousValTemp); i <= max(currentVal, previousValTemp); i++ {
			if i%100 == 0 {
				// fmt.Println("hits zero at", i)
				zeroCounts += 1
			}
		}

		// Handle wraparound
		numRotations := 0 // Number of times that the currentVal is wrapped past 100
		polarity := 1     // 1 if $currentVal is positive, -1 is $currentVal is negative
		if currentVal >= 100 || currentVal < 0 {
			// Gets the specific number of times that currentVal is wrapped past 100
			numRotations = utils.AbsInt(currentVal) / 100
			// Gets the Absolute Int number of rotations
			numRotations = max(utils.AbsInt(numRotations), 1)
			// Sets $negative to -1 if the $currentVal is below zero
			if currentVal < 0 {
				polarity = -1
			}
			// Subtracts 100 * the number of rotations * the polarity (negative/positive)
			currentVal -= (100 * numRotations) * polarity
		}

		previousVal = currentVal

	}
	return zeroCounts
}

func main() {
	fmt.Println("Result A:", partA())
	fmt.Println("Result B:", partB())
}
