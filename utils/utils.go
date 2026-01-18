package utils

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
}

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func ReadLines(path string) ([]string, error) {
	// Reads the lines of a file as individual items in a slice and
	// returns them, then closes the file.
	file, err := os.Open(path)
	Check(err)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

func IntFloor(number int) int {
	// Takes an integer, floors it, and returns an int
	return int(math.Floor(float64(number)))
}

func FlattenSlice(inputSlice []string) (outputString string) {
	// Flattens a slice into a single string
	outputString = strings.Join(inputSlice, "")

	return outputString
}

func BoxFilter(inputSlice []string, xRadius int, yRadius int, x int, y int) []string {
	// Safe box filter given an input slice of Y (slice, vertical) and X (string, horizontal)
	// $xRadius for string search radius, $yRadius for slice search radius
	// Returns a slice with the box filter's search radius
	var outputSlice []string
	yMin, yMax := max(y-yRadius, 0), min(y+yRadius+1, len(inputSlice)) // Safe y min and y max
	for _, Y := range inputSlice[yMin:yMax] {
		xMin, xMax := max(x-xRadius, 0), min(x+xRadius+1, len(Y)) // safe x min and x max
		outputSlice = append(outputSlice, Y[xMin:xMax])
	}

	return outputSlice
}
