package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"strings"
)

type MinMaxRange struct {
	Min int
	Max int
}

type XYZ struct {
	X int
	Y int
	Z int
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

func CountAll(inputSlice []bool) (outputCount int, err error) {
	for _, item := range inputSlice {
		if item == true {
			outputCount++
		}
	}
	return outputCount, nil
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

func AOCFileReadToSlice(test bool, day int) (output []string, err error) {
	// Basic file importing
	ex, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	folder := "input"

	if test {
		folder = "testinput"
	}

	dayStr := fmt.Sprintf("%02d", day)

	path := filepath.Join(ex, folder, dayStr+".txt")
	data, err := ReadLines(path)
	if err != nil {
		return nil, err
	}

	return data, nil
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

func RangeOverlapsSorted(rangeA MinMaxRange, rangeB MinMaxRange) (overlaps bool, err error) {
	// Takes two given inputs of $minMaxRange, with $rangeB having a higher starting int,
	// 	checks if they overlap, and returns true if they do overlap, and false if they don't
	return (rangeB.Min < rangeA.Max && rangeA.Min < rangeB.Max), nil
}

func CleanOverlappingRanges(inputRanges []MinMaxRange) (outputRanges []MinMaxRange, err error) {
	// Takes a sorted input slice of minMaxRange{} items, checks for any overlapping min/max ranges in the slice
	// 	and if there are any overlaps, removes the offending items and returns a clean range instead
	outputRanges = append(outputRanges, inputRanges[0])
	for i := 1; i <= len(inputRanges)-1; i++ {
		last := &outputRanges[len(outputRanges)-1]
		curr := inputRanges[i]
		if curr.Min <= last.Max {
			last.Max = max(last.Max, curr.Max)
		} else {
			outputRanges = append(outputRanges, curr)
		}
	}

	return outputRanges, err
}

func SplitStrIntoArbitraryLength(inputStr string, endStringsLength int) (outputStrAsSlice []string, err error) {
	// Takes an input string $inputStr, and an input int $endStringsLength,
	// 	and outputs a slice of strings of $endStringsLength each
	count := 0
	currentStr := bytes.Buffer{}
	for i := range len(inputStr) {
		currentStr.WriteString(inputStr[i : i+1])
		count++
		if count == endStringsLength {
			outputStrAsSlice = append(outputStrAsSlice, currentStr.String())
			currentStr = bytes.Buffer{}
			count = 0
		}
	}
	if !(currentStr.String() == "") {
		outputStrAsSlice = append(outputStrAsSlice, currentStr.String())
	}
	return outputStrAsSlice, nil
}

// Takes two inputs of type XYZ, and returns the distance between them
func XYZEuclidianDistance(inputA XYZ, inputB XYZ) (distance float64, err error) {
	xDist := math.Pow(float64(inputA.X-inputB.X), 2)
	yDist := math.Pow(float64(inputA.Y-inputB.Y), 2)
	zDist := math.Pow(float64(inputA.Z-inputB.Z), 2)
	distance = math.Sqrt(xDist + yDist + zDist)
	return distance, nil
}
