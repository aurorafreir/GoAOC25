package main

import (
	"fmt"
	"os"
	"path/filepath"

	"example.com/utils"
)

func partA() int {
	// Basic file importing
	ex, err := os.Getwd()
	utils.Check(err)

	path := filepath.Join(ex, "input.txt")
	data, err := utils.ReadLines(path)
	utils.Check(err)
	for _, line := range data {

	}

	return 0
}

func partB() int {
	return 0
}

func main() {
	fmt.Println("Result A:", partA())
	fmt.Println("Result B:", partB())
}
