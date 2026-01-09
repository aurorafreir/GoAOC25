package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strconv"

	"example.com/generics"
)

func partA() int {
	currentVal := 50
	zeroCounts := 0

	ex, err := os.Getwd()
	generics.Check(err)

	path := filepath.Join(ex, "input.txt")
	data, err := generics.ReadLines(path)
	generics.Check(err)
	for _, line := range data {
		direction := line[0:1]
		num := line[1:]
		intconv, err := strconv.Atoi(num)
		generics.Check(err)

		switch direction {
		case "L":
			currentVal -= intconv
		case "R":
			currentVal += intconv
		}

		if currentVal%100 == 0 {
			zeroCounts += 1
		}

	}
	return zeroCounts
}

func intFloor(number int) int {
	// Takes an integer, floors it, and returns an int
	return int(math.Floor(float64(number)))
}

func partB() int {
	currentVal := 50
	zeroCounts := 0

	ex, err := os.Getwd()
	generics.Check(err)

	path := filepath.Join(ex, "testinput.txt")
	data, err := generics.ReadLines(path)
	generics.Check(err)
	for _, line := range data {

		// Split the line (e.g. "L68") into $direction ("L") and $strnum ("68"),
		// then converts $num to an integer
		direction := line[0:1]
		strnum := line[1:]
		num, err := strconv.Atoi(strnum)
		generics.Check(err)

		// Add or subtract the new number from $currentVal
		switch direction {
		case "L":
			currentVal -= num
		case "R":
			currentVal += num
		}

		// Handle wraparound
		numRotations := 0 // Number of times that the currentVal is wrapped past 100
		polarity := 1     // 1 if $currentVal is positive, -1 is $currentVal is negative
		if currentVal >= 100 || currentVal < 0 {
			// Gets the specific number of times that currentVal is wrapped past 100
			numRotations = generics.AbsInt(currentVal) / 100
			// Gets the Absolute Int number of rotations
			numRotations = max(generics.AbsInt(numRotations), 1)
			// Sets $negative to -1 if the $currentVal is below zero
			if currentVal < 0 {
				polarity = -1
			}
			// Subtracts 100 * the number of rotations * the polarity (negative/positive)
			currentVal -= (100 * numRotations) * polarity
		}
		fmt.Println(currentVal, numRotations)

	}
	return zeroCounts
}

func main() {
	fmt.Println("Result A:", partA())
	fmt.Println("Result B:", partB())
}
