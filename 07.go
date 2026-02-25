package main

import (
	"fmt"

	"example.com/utils"
)

func d7p1() (output int, err error) {
	testing := false
	data, err := utils.AOCFileReadToSlice(testing, 7)
	utils.Check(err)

	startingIndex := 0
	for index, char := range data[0] {
		if string(char) == "S" {
			startingIndex = index
		}
	}

	if testing {
		fmt.Println(startingIndex)
	}

	tachyons := make([][]bool, len(data))
	tachyons[0] = make([]bool, len(data[0]))
	tachyons[0][startingIndex] = true

	tachyonSplits := 0

	for indexY := 1; indexY < len(data); indexY++ {
		tachyons[indexY] = make([]bool, len(data[indexY]))
		if testing {
			fmt.Println(string(data[indexY]))
		}
		for indexX, char := range data[indexY] {
			if string(char) == "^" && tachyons[indexY-1][indexX] == true {
				tachyons[indexY][max(indexX-1, 0)] = true
				tachyons[indexY][min(indexX+1, len(data[indexY]))] = true
				tachyonSplits++
			} else if tachyons[indexY-1][indexX] == true {
				tachyons[indexY][indexX] = true
			}
		}
	}

	if testing {
		fmt.Println(tachyons)
	}

	output = tachyonSplits
	return output, nil
}

func tachyonSplitter(inputSlice []string, tachyons [][]bool, startingYIndex int, startingXIndex int, direction bool, splits int) (splitsOut int) {
	if tachyons[startingYIndex][startingXIndex] {
		return splits
	}
	splits++
	for indexY := startingYIndex; indexY < len(inputSlice); indexY++ {
		for indexX := startingXIndex; indexX < len(inputSlice[indexY]); indexX++ {
			if string(inputSlice[indexY][indexX]) == "^" && tachyons[indexY-1][indexX] == true {
				if direction {
					fmt.Println("right")
					tachyons[indexY][min(indexX+1, len(inputSlice[indexY]))] = true
					splits = tachyonSplitter(inputSlice, tachyons, indexY, indexX, false, splits)
					// splits++
				} else {
					fmt.Println("left")
					tachyons[indexY][max(indexX-1, 0)] = true
					splits = tachyonSplitter(inputSlice, tachyons, indexY, indexX+1, true, splits)
					// splits++
				}
			} else if tachyons[indexY-1][indexX] == true {
				tachyons[indexY][indexX] = true
			}
		}
		fmt.Println(splits, indexY, tachyons[indexY])
	}

	return splits
}

func d7p2() (output int, err error) {
	testing := true
	data, err := utils.AOCFileReadToSlice(testing, 7)
	utils.Check(err)

	startingIndex := 0
	for index, char := range data[0] {
		if string(char) == "S" {
			startingIndex = index
		}
	}

	// initiate tachyons 2d slice
	tachyons := make([][]bool, len(data))
	for index := range data {
		tachyons[index] = make([]bool, len(data[index]))
	}
	tachyons[0][startingIndex] = true

	splits := tachyonSplitter(data, tachyons, 1, 0, true, 0)

	fmt.Println("splits:", splits)

	return output, nil
}

// func d7p2() (output int, err error) {
// 	testing := true
// 	data, err := utils.AOCFileReadToSlice(testing, 7)
// 	utils.Check(err)

// 	startingIndex := 0
// 	for index, char := range data[0] {
// 		if string(char) == "S" {
// 			startingIndex = index
// 		}
// 	}

// 	if testing {
// 		fmt.Println(startingIndex)
// 	}

// 	tachyons := make([][]bool, len(data))
// 	tachyons[0] = make([]bool, len(data[0]))
// 	tachyons[0][startingIndex] = true

// 	tachyonSplits := 0

// 	for indexY := 1; indexY < len(data); indexY++ {
// 		tachyons[indexY] = make([]bool, len(data[indexY]))
// 		if testing {
// 			fmt.Println(string(data[indexY]))
// 		}
// 		for indexX, char := range data[indexY] {
// 			if string(char) == "^" && tachyons[indexY-1][indexX] == true {
// 				tachyons[indexY][max(indexX-1, 0)] = true
// 				tachyons[indexY][min(indexX+1, len(data[indexY]))] = true
// 				tachyonSplits++
// 			} else if tachyons[indexY-1][indexX] == true {
// 				tachyons[indexY][indexX] = true
// 			}
// 		}
// 	}

// 	// for _, line := range tachyons {
// 	// 	outString := ""
// 	// 	for _, i := range line {
// 	// 		if i {
// 	// 			outString = outString + strings.ReplaceAll(strconv.FormatBool(i), "true", "|")
// 	// 		} else {
// 	// 			outString = outString + strings.ReplaceAll(strconv.FormatBool(i), "false", ".")
// 	// 		}
// 	// 	}
// 	// 	fmt.Println(outString)
// 	// }

// 	output = tachyonSplits
// 	return output, nil
// }

func main() {
	// fmt.Println(d7p1())
	fmt.Println(d7p2())
}
