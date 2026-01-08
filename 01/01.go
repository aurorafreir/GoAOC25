package main

import (
	"fmt"
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

func partB() int {
	// previousVal := 50
	currentVal := 50
	zeroCounts := 0

	ex, err := os.Getwd()
	generics.Check(err)

	path := filepath.Join(ex, "input.txt")
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

		// previousVal = currentVal

	}
	return zeroCounts
}

func main() {
	fmt.Println("Result A:", partA())
	fmt.Println("Result B:", partB())
}
