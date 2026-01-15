package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"example.com/utils"
)

/*
The batteries are arranged into banks; each line of digits in your input
corresponds to a single bank of batteries. Within each bank, you need to
turn on exactly two batteries; the joltage that the bank produces is equal
to the number formed by the digits on the batteries you've turned on.
For example, if you have a bank like 12345 and you turn on batteries 2 and 4,
the bank would produce 24 jolts. (You cannot rearrange batteries.)

You'll need to find the largest possible joltage each bank can produce.
In the above example:

    In 987654321111111, you can make the largest joltage possible, 98, by turning on the first two batteries.
    In 811111111111119, you can make the largest joltage possible by turning on the batteries labeled 8 and 9, producing 89 jolts.
    In 234234234234278, you can make 78 by turning on the last two batteries (marked 7 and 8).
    In 818181911112111, the largest joltage you can produce is 92.

The total output joltage is the sum of the maximum joltage from each bank,
so in this example, the total output joltage is 98 + 89 + 78 + 92 = 357.
*/

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
			for y := x + 1; y < lenOfNum; y++ {
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

func findHighestNumInSlicedStr(inputStr string, startIntIncl int, endIntExcl int) (int, int, error) {
	highestInt := 0
	highestIntIndex := startIntIncl

	// Loop through the $inputStr from $startIntIncl to $endIntExcl, convert to an int, and error check
	for index := startIntIncl; index <= endIntExcl; index++ {
		indexItemStr := inputStr[index : index+1]
		intItem, err := strconv.Atoi(indexItemStr)
		utils.Check(err)

		// No point in checking further if item is 9, return next index
		if intItem == 9 {
			return intItem, index + 1, nil
		}
		// Find highest item and the next index
		if intItem > highestInt {
			highestInt = intItem
			highestIntIndex = index + 1
		}
	}

	return highestInt, highestIntIndex, nil
}

func partB() (int, error) {
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
		// Sliding window code
		endIntLen := 12
		startRange := 0
		endRange := lenOfNum - endIntLen
		endStrSlice := []string{}
		endStr := ""
		for range endIntLen {
			highestInt, highestIntIndex, err := findHighestNumInSlicedStr(line, startRange, endRange)
			utils.Check(err)
			startRange = highestIntIndex
			endRange = lenOfNum - (endIntLen - len(endStrSlice)) + 1 // Moves $endRange forward
			endStrSlice = append(endStrSlice, strconv.Itoa(highestInt))
			endStr = strings.Join(endStrSlice, "")
		}
		maxInt, err = strconv.Atoi(endStr)

		maxInts = append(maxInts, maxInt)
	}

	// Addition of each int in $maxInts
	for _, int := range maxInts {
		joltage += int
	}

	return joltage, nil
}

func main() {
	// fmt.Println("Result A:", partA())

	funcTimeStartA := time.Now()
	pA := partA()
	fmt.Println("Result A:", pA)
	fmt.Println("Function A took:", time.Since(funcTimeStartA))

	funcTimeStartB := time.Now()
	pB, err := partB()
	utils.Check(err)
	fmt.Println("Result B:", pB)
	fmt.Println("Function B took:", time.Since(funcTimeStartB))
}
