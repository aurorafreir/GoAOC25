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

func d7p2() (output int, err error) {

	return output, nil
}
