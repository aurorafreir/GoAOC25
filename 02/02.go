package main

import (
	"fmt"
	"os"
	"path/filepath"

	"example.com/generics"
)

func partA() int {
	ex, err := os.Getwd()
	generics.Check(err)

	path := filepath.Join(ex, "testinput.txt")
	data, err := generics.ReadLines(path)
	generics.Check(err)

	singleString := data[0]
	fmt.Println(singleString)

	return 0
}

func partB() int {
	return 0
}

func main() {
	fmt.Println("Result A:", partA())
	fmt.Println("Result B:", partB())
}
