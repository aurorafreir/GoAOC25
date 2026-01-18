package main

import (
	"strconv"
	"strings"

	"example.com/utils"
)

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

func findForwardHighestInt(inputData []string, resultIntLength int) []int {
	var maxInts []int
	for _, line := range inputData {
		lenOfNum := len(line)
		maxInt := 0
		// Sliding window code
		endIntLen := resultIntLength
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
		maxInt, err := strconv.Atoi(endStr)
		utils.Check(err)

		maxInts = append(maxInts, maxInt)
	}

	return maxInts
}

func d3(maxIntLength int) (int, error) {
	joltage := 0 // i like this word

	data, err := utils.AOCFileReadToSlice(false, 3)
	utils.Check(err)

	maxInts := findForwardHighestInt(data, maxIntLength)

	// Addition of each int in $maxInts
	for _, int := range maxInts {
		joltage += int
	}

	return joltage, nil
}

// func main() {
// 	funcTimeStartA := time.Now()
// 	pA, err := day3Handler(2)
// 	utils.Check(err)
// 	fmt.Println("D3p1 result:", pA)
// 	fmt.Println("D3p1 took:", time.Since(funcTimeStartA).Microseconds(), "microseconds")

// 	funcTimeStartB := time.Now()
// 	pB, err := day3Handler(12)
// 	utils.Check(err)
// 	fmt.Println("D3p2 result:", pB)
// 	fmt.Println("D3p2 took:", time.Since(funcTimeStartB).Microseconds(), "microseconds")
// }
