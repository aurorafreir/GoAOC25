package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"example.com/utils"
)

func partA() int {
	var maxInts []int
	joltage := 0 // i like this word

	// Basic file importing
	ex, err := os.Getwd()
	utils.Check(err)

	path := filepath.Join(ex, "input.txt")
	data, err := utils.ReadLines(path)
	utils.Check(err)

	for _, line := range data {
		lenOfNum := len(line)
		maxInt := 0
		// Loop through each possible forward-only iteration of the input line
		for x := 0; x <= lenOfNum; x++ {
			for y := x; y < lenOfNum; y++ {
				if y == x {
					continue
				}
				// Concatenate item $X and item $Y as a string
				var b bytes.Buffer
				b.WriteString(line[x : x+1])
				b.WriteString(line[y : y+1])
				// Convert to integer
				asInt, err := strconv.Atoi(b.String())
				utils.Check(err)
				// Replace $maxInt if $asInt is higher
				if asInt > maxInt {
					maxInt = asInt
				}
			}
		}
		maxInts = append(maxInts, maxInt)
	}

	// Addition of each int in $maxInts
	for _, int := range maxInts {
		joltage += int
	}
	return joltage
}

func partB() int {
	return 0
}

func main() {
	fmt.Println("Result A:", partA())
	fmt.Println("Result B:", partB())
}
