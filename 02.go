package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"example.com/utils"
)

func lenItoa(i int) int {
	// From stackoverflow https://stackoverflow.com/a/68122831
	return len(strconv.Itoa(i))
}

func d2p1() (output int, err error) {
	data, err := utils.AOCFileReadToSlice(false, 2)
	utils.Check(err)

	// Get just the first line as a string
	singleString := data[0]

	sliceOfRangeStrings := strings.Split(singleString, ",") // Split single string into each range as a string
	sliceOfRangeInts := []utils.MinMaxRange{}

	for _, item := range sliceOfRangeStrings {
		// Get each range as a slice of two strings, then convert to ints,
		// 		then append to $sliceOfRangeInts as a $MinMaxRange struct
		ranges := strings.Split(item, "-")
		rangeMin, _ := strconv.Atoi(ranges[0])
		rangeMin = max(rangeMin, 10) // No point in checking integers under 10
		rangeMax, _ := strconv.Atoi(ranges[1])
		sliceOfRangeInts = append(sliceOfRangeInts, utils.MinMaxRange{rangeMin, rangeMax})
	}

	invalidCount := 0

	// Find all the invalid items in the ranges
	for _, ranges := range sliceOfRangeInts {
		for number := ranges.Min; number <= ranges.Max; number++ {
			numberAsStr := strconv.Itoa(number)
			halfLength := len(numberAsStr) / 2
			if numberAsStr[0:halfLength] == numberAsStr[halfLength:] { // Split the item into two same length items and see if they match
				invalidCount += number
			}
		}
	}

	return invalidCount, nil
}

func d2p2() (output int, err error) {
	testing := false
	data, err := utils.AOCFileReadToSlice(testing, 2)
	utils.Check(err)

	// Get just the first line as a string
	singleString := data[0]

	sliceOfRangeStrings := strings.Split(singleString, ",") // Split single string into each range as a string
	sliceOfRangeInts := []utils.MinMaxRange{}

	for _, item := range sliceOfRangeStrings {
		// Get each range as a slice of two strings, then convert to ints,
		// 		then append to $sliceOfRangeInts as a $MinMaxRange struct
		ranges := strings.Split(item, "-")
		rangeMin, _ := strconv.Atoi(ranges[0])
		rangeMin = max(rangeMin, 10) // No point in checking integers under 10
		rangeMax, _ := strconv.Atoi(ranges[1])
		sliceOfRangeInts = append(sliceOfRangeInts, utils.MinMaxRange{rangeMin, rangeMax})
	}

	invalidCount := 0

	// Find all the invalid items in the ranges
	for _, ranges := range sliceOfRangeInts {
		for number := ranges.Min; number <= ranges.Max; number++ {
			numberAsStr := strconv.Itoa(number)
			halfLength := len(numberAsStr) / 2
			cleanMults := []int{}
			for length := 1; length <= halfLength; length++ {
				if len(numberAsStr)%length == 0 {
					cleanMults = append(cleanMults, length)
				}
			}
			if testing {
				fmt.Println(numberAsStr, "cleanMults", cleanMults)
			}
		out:
			for _, mult := range cleanMults {
				splitString, _ := utils.SplitStrIntoArbitraryLength(numberAsStr, mult)
				allTheSame := !slices.ContainsFunc(splitString, func(s string) bool { return s != splitString[0] })
				if testing {
					fmt.Println(allTheSame)
				}
				if allTheSame {
					invalidCount += number
					break out
				}
			}
		}
	}

	return invalidCount, nil
}
