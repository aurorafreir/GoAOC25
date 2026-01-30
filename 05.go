package main

import (
	"sort"
	"strconv"
	"strings"

	"example.com/utils"
)

func returnRangesAndIDs(data []string, onlyRanges bool) ([]utils.MinMaxRange, []int) {
	pastEmptyLine := false

	ranges := make([]utils.MinMaxRange, 0)
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
			rngInt := utils.MinMaxRange{rngStartInt, rngEndInt}
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
			if startEndRange.Min <= id && id <= startEndRange.Max {
				freshItems[id] = true
			}
		}
	}

	return len(freshItems), nil
}

func d5p2() (int, error) {
	data, err := utils.AOCFileReadToSlice(false, 5)
	utils.Check(err)

	count := 0
	inRanges := []int{}
	inOutRangesMap := make(map[int]int)
	sortedRanges := []utils.MinMaxRange{}
	cleanRanges := []utils.MinMaxRange{}

	ranges, _ := returnRangesAndIDs(data, true)

	// This section of sorting the data by input range can almost CERTAINLY be done in less
	// 		lines of code, might revisit at some point
	// Map wih input ranges as key and output ranges as max
	for _, startEndRange := range ranges {
		if _, exists := inOutRangesMap[startEndRange.Min]; exists { // Checks for collisions and sets value to highest Max range
			inOutRangesMap[startEndRange.Min] = max(startEndRange.Max, inOutRangesMap[startEndRange.Min])
			continue
		} else {
			inOutRangesMap[startEndRange.Min] = startEndRange.Max
		}
	}

	// Slice of just input ranges to be sorted
	for _, startEndRange := range ranges {
		inRanges = append(inRanges, startEndRange.Min)
	}
	sort.Ints(inRanges) // sorting happens in place?? weird.

	// Spit out a clean set of ranges based on the sorted input keys
	for _, item := range inRanges {
		sortedRanges = append(sortedRanges, utils.MinMaxRange{item, inOutRangesMap[item]})
	}

	cleanRanges, _ = utils.CleanOverlappingRanges(sortedRanges)

	// Inclusive range addition
	for _, i := range cleanRanges {
		count += (i.Max + 1 - i.Min)
	}

	return count, nil
}
