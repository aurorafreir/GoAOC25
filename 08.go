package main

import (
	"errors"
	"fmt"
	"slices"
	"sort"
	"strconv"
	"strings"

	"example.com/utils"
)

func d8p1() (output int, err error) {
	testing := true
	data, err := utils.AOCFileReadToSlice(testing, 8)
	utils.Check(err)

	// Convert lines to utils.XYZ{}
	xyzPositions := []utils.XYZ{}
	for _, line := range data {
		xyz := utils.XYZ{}
		xyzSlice := strings.Split(line, ",")
		xyz.X, _ = strconv.Atoi(xyzSlice[0])
		xyz.Y, _ = strconv.Atoi(xyzSlice[1])
		xyz.Z, _ = strconv.Atoi(xyzSlice[2])
		xyzPositions = append(xyzPositions, xyz)
	}

	if testing {
		// fmt.Println(xyzPositions)
	}

	xyzDistancesAndValues := map[float64][]utils.XYZ{}
	// map, sort, clip
	for indexX := 0; indexX < len(xyzPositions); indexX++ {
		for indexY := indexX + 1; indexY < len(xyzPositions); indexY++ {
			if !(indexX == indexY) {
				dist, _ := utils.XYZEuclidianDistance(xyzPositions[indexX], xyzPositions[indexY])
				_, ok := xyzDistancesAndValues[dist]
				if ok {
					fmt.Println("Well that wasn't supposed to happen :/")
					return 0, errors.New("Error: xyzDistancesAndValues collision")
				}
				xyzDistancesAndValues[dist] = []utils.XYZ{xyzPositions[indexX], xyzPositions[indexY]}
			}
		}
	}

	xyzDistanceKeys := []float64{}
	for key := range xyzDistancesAndValues {
		xyzDistanceKeys = append(xyzDistanceKeys, key)
	}
	sort.Float64s(xyzDistanceKeys)

	circuitSets := [][]utils.XYZ{}
	// we can assume the first two items aren't in the list
	circuitSets = append(circuitSets, []utils.XYZ{xyzDistancesAndValues[xyzDistanceKeys[0]][0], xyzDistancesAndValues[xyzDistanceKeys[0]][1]})

	// Check if next item is in $circuitSets, if it is then move to next slice, if it isn't then append to current slice
	// currentSlice := 0
	// for _, item := range xyzDistanceKeys[1:] {
	for i := 0; i < len(xyzDistanceKeys); i++ {
		xyzA, xyzB := xyzDistancesAndValues[xyzDistanceKeys[i]][0], xyzDistancesAndValues[xyzDistanceKeys[i]][1]
		fmt.Println(xyzA, xyzB)
		// out:
		for _, circuitSlice := range circuitSets {
			if !(slices.Contains(circuitSlice, xyzA) && slices.Contains(circuitSlice, xyzB)) {
				if !slices.Contains(circuitSlice, xyzA) {
					circuitSlice = append(circuitSlice, xyzA)
				} else {
					circuitSlice = append(circuitSlice, xyzB)
				}
				// i++
			} else {
				circuitSets = append(circuitSets, []utils.XYZ{xyzA, xyzB})
				// i++
				// break out
			}
		}
	}

	for index, circuitSet := range circuitSets {
		fmt.Println(index, circuitSet)
	}

	return 0, nil
}

// func main() {
// 	fmt.Println(d8p1())
// }
