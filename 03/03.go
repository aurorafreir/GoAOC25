package main

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
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

func combineAndReturnIntFromStr(line string, startInt int) (int, error) {
	var byteStr bytes.Buffer
	byteStr.WriteString(line[startInt : startInt+12])
	// Convert to integer
	asInt, err := strconv.Atoi(byteStr.String())
	fmt.Println(asInt)
	return asInt, err
}

func partB() (int, error) {
	// unworking currently

	var maxInts []int
	joltage := 0 // i like this word

	// Basic file importing
	ex, err := os.Getwd()
	utils.Check(err)

	path := filepath.Join(ex, "testinput.txt")
	data, err := utils.ReadLines(path)
	utils.Check(err)

	for lineNum, line := range data {
		timeStart := time.Now()
		lenOfNum := len(line)
		maxInt := 0
		fmt.Println(lineNum+1, "out of:", len(data))
		// Get clean direct forward 12 length long numbers
		for a := 0; a <= lenOfNum-12; a++ {
			asInt, err := combineAndReturnIntFromStr(line, a)
			utils.Check(err)
			// Replace $maxInt if $asInt is higher
			if asInt > maxInt {
				maxInt = asInt
			}
		}

		maxInts = append(maxInts, maxInt)
		fmt.Println(time.Since(timeStart))
	}
	// Addition of each int in $maxInts
	for _, int := range maxInts {
		joltage += int
	}

	return joltage, nil
}

func main() {
	// fmt.Println("Result A:", partA())

	// funcTimeStart := time.Now()
	// pB, err := partB()
	// utils.Check(err)
	// fmt.Println("Result B:", pB)
	// fmt.Println("Function took:", time.Since(funcTimeStart))

	// lenOfOutputNums := 6
	items := []string{"120240332040", "120240332040"}
	for index, item := range items {
		fmt.Println(index, item)
		// highestNumStr := "000000000000"
		for i := 0; i < len(item); i++ {
			fmt.Println(item[i : i+1])
		}
	}
}
