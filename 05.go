package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"example.com/utils"
)

func returnRangesAndIDs(data []string, onlyRanges bool) ([][]int, []int) {
	pastEmptyLine := false

	ranges := make([][]int, 0)
	var ids []int

	for _, line := range data {
		if line == "" {
			pastEmptyLine = true
		}
		if pastEmptyLine && onlyRanges {
			return ranges, ids
		}
		if pastEmptyLine {
			intConv, _ := strconv.Atoi(line)
			ids = append(ids, intConv)
		} else if line != "" {
			rng := strings.Split(line, "-")
			rngStartInt, _ := strconv.Atoi(rng[0])
			rngEndInt, _ := strconv.Atoi(rng[1])
			rngInt := []int{rngStartInt, rngEndInt}
			ranges = append(ranges, rngInt)
		}
	}
	return ranges, ids
}

func d5p1() (int, error) {
	data, err := utils.AOCFileReadToSlice(false, 5)
	utils.Check(err)

	ranges, ids := returnRangesAndIDs(data, false)

	freshItems := make(map[int]bool)

	for _, id := range ids {
		for _, startEndRange := range ranges {
			if startEndRange[0] <= id && id <= startEndRange[1] {
				freshItems[id] = true
			}
		}
	}

	return len(freshItems), nil
}

func d5p2() (int, error) {
	data, err := utils.AOCFileReadToSlice(true, 5)
	utils.Check(err)

	// [x] Get ranges
	// [x] Get a slice of just the input range numbers
	// [x] Sort input range sort
	// [x] Create map of input ranges, key being start range and value being output range
	// [x] Loop through sorted slice to grab k,v from map and export a sorted input/output slice
	// [x] loop through sorted slice, and given it's output range, see if the next input range is
	// 			lower than the current output range, if it is, append [$current[input], $next[ouxtput]]
	// [ ] For each item in the sorted ranges, find out if the output range is bigger than the next input range.
	// 			If it isn't then output it as is. If it is bigger than the next input range then loop through the
	// 			future inOutRanges and find the next one that is bigger than the current outRange, then append
	// 			the current inputRange and the previous future outputRange, then skip to the next inputRange

	// count := 0
	var inRanges []int
	inOutRangesMap := make(map[int]int)
	sortedRanges := [][]int{}
	cleanRanges := [][]int{}

	ranges, _ := returnRangesAndIDs(data, true)

	for _, startEndRange := range ranges {
		inOutRangesMap[startEndRange[0]] = startEndRange[1]
	}

	for _, startEndRange := range ranges {
		inRanges = append(inRanges, startEndRange[0])
	}

	sort.Ints(inRanges) // this does it in place?? weird.

	for _, item := range inRanges {
		newItem := []int{item, inOutRangesMap[item]}
		sortedRanges = append(sortedRanges, newItem)
	}

	fmt.Println(inOutRangesMap)
	fmt.Println("only input ranges:", inRanges)
	fmt.Println("sorted range slice:", sortedRanges)

	currentPos := 0
	lenOfSortedRange := len(sortedRanges)
	for range lenOfSortedRange {
		// Exit loop cleanly
		if currentPos >= lenOfSortedRange-1 {
			return 0, nil
		}

		fmt.Println(sortedRanges[currentPos])

		// Clean output for current item, no overlapping ranges, append with no changes
		if sortedRanges[currentPos][1] < sortedRanges[currentPos+1][0] {
			fmt.Println("clean output, no overlaps:", sortedRanges[currentPos])
			cleanRanges = append(cleanRanges, []int{sortedRanges[currentPos][0], sortedRanges[currentPos][1]})
			currentPos++
			continue
		}

		// Check if the current next input range is lower than the current highest output
		currentInputRange := sortedRanges[currentPos][0]
		currentHighestOutputRange := sortedRanges[currentPos][1]

		for indexY := currentPos + 1; indexY < lenOfSortedRange-1; indexY++ {
			fmt.Println(indexY, sortedRanges[indexY], currentHighestOutputRange)
			// currentHighestOutputRange = max(currentHighestOutputRange, sortedRanges[indexY][1])
			if sortedRanges[indexY+1][0] > currentHighestOutputRange {
				currentHighestOutputRange = sortedRanges[indexY][1]
				currentPos = indexY + 1
				continue

			}
		}

		// Append
		cleanRanges = append(cleanRanges, []int{currentInputRange, currentHighestOutputRange})

		fmt.Println("clean ranges:", cleanRanges)

	}

	return 0, nil
}
