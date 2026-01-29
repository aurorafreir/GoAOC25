package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"example.com/utils"
)

type minMaxRange struct {
	min int
	max int
}

func returnRangesAndIDs(data []string, onlyRanges bool) ([]minMaxRange, []int) {
	pastEmptyLine := false

	ranges := make([]minMaxRange, 0)
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
			rngInt := minMaxRange{rngStartInt, rngEndInt}
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
			if startEndRange.min <= id && id <= startEndRange.max {
				freshItems[id] = true
			}
		}
	}

	return len(freshItems), nil
}

func rangeOverlaps(rangeA minMaxRange, rangeB minMaxRange) (overlaps bool) {
	return (rangeB.min < rangeA.max && rangeA.min < rangeB.max)
}

func cleanOverlappingRanges(inputRanges []minMaxRange) (outputRanges []minMaxRange, err error) {
	// Takes a sorted input slice of minMaxRange{} items, checks for any overlapping min/max ranges in the slice
	// 	and if there are any overlaps, removes the offending items and returns a clean range instead

	outputRanges = append(outputRanges, inputRanges[0])
	for i := 1; i <= len(inputRanges)-1; i++ {
		last := &outputRanges[len(outputRanges)-1]
		curr := inputRanges[i]
		if curr.min <= last.max {
			last.max = max(last.max, curr.max)
		} else {
			outputRanges = append(outputRanges, curr)
		}
	}

	return outputRanges, err
}

func d5p2() (int, error) {
	data, err := utils.AOCFileReadToSlice(false, 5)
	utils.Check(err)

	count := 0
	inRanges := []int{}
	inOutRangesMap := make(map[int]int)
	sortedRanges := []minMaxRange{}
	cleanRanges := []minMaxRange{}

	ranges, _ := returnRangesAndIDs(data, true)

	// This section of sorting the data by input range can almost CERTAINLY be done in less
	// 		lines of code, might revisit at some point
	// Map wih input ranges as key and output ranges as max
	for _, startEndRange := range ranges {
		inOutRangesMap[startEndRange.min] = startEndRange.max
	}

	// Slice of just input ranges
	for _, startEndRange := range ranges {
		inRanges = append(inRanges, startEndRange.min)
	}

	sort.Ints(inRanges) // sorting happens in place?? weird.

	// Spit out a clean set of ranges based on the sorted input keys
	for _, item := range inRanges {
		sortedRanges = append(sortedRanges, minMaxRange{item, inOutRangesMap[item]})
	}

	// fmt.Println(inOutRangesMap)
	// fmt.Println("only input ranges:", inRanges)
	// fmt.Println("sorted range slice:", sortedRanges)

	cleanRanges, err = cleanOverlappingRanges(sortedRanges)
	utils.Check(err)
	fmt.Println("clean ranges:", cleanRanges)
	for _, i := range cleanRanges {
		fmt.Println(i.max + 1 - i.min)
		count += (i.max + 1 - i.min)
		// fmt.Println(i)
	}

	return count, nil
}

func main() {
	fmt.Println(d5p2())
}
