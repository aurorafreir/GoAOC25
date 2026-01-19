package main

import (
	"fmt"
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

	ranges, _ := returnRangesAndIDs(data, true)

	sortedRanges := make(map[int][][]int)

	cleanRanges := make([][]int, 0)

	count := 0

	for index, startEndRange := range ranges[:len(ranges)-1] {
		if startEndRange[1] >= ranges[index+1][0] {
			fmt.Println(startEndRange, ranges[index+1][0])
		}
	}

	fmt.Println(ranges, cleanRanges)

	return count, nil
}
