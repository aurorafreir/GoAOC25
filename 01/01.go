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
		//fmt.Println(direction, intconv)

		switch direction {
		case "L":
			currentVal -= intconv
		case "R":
			currentVal += intconv
		}

		if currentVal%100 == 0 {
			zeroCounts += 1
		}
		//fmt.Println(currentVal)

	}
	//fmt.Println("Result:", zeroCounts)
	return zeroCounts
}

func partB() int {
	previousVal := 50
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

		if intconv > 100 {
			numRotations := 0
			numRotations = (generics.AbsInt(currentVal) / 100) >> 0
			zeroCounts += numRotations
		}

		if currentVal >= 0 && currentVal <= 99 {

		} else if currentVal > 99 {
			numRotations := 0
			if currentVal == 100 {
				numRotations = 1
			} else {
				numRotations = (currentVal / 100) >> 0
				numRotations = max(generics.AbsInt(numRotations), 1)
			}

			currentVal -= (100 * numRotations)
		} else if currentVal < 0 {
			numRotations := (currentVal / 100) >> 0
			numRotations = max(generics.AbsInt(numRotations), 1)
			currentVal += (100 * numRotations)
		}

		if currentVal == 0 {
			zeroCounts += 1
		} else if currentVal > 100 || currentVal < 0 {
			for i := min(previousVal, currentVal); i <= max(previousVal, currentVal); i++ {
				if i == 0 {
					zeroCounts += 1
				}
			}
		}

		previousVal = currentVal

	}
	return zeroCounts
}

func main() {
	fmt.Println("Result A:", partA())
	fmt.Println("Result B:", partB())
}
