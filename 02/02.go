package main

import (
	"fmt"
	"os"
	"path/filepath"

	"example.com/generics"
)

func partA() string {
	ex, err := os.Getwd()
	generics.Check(err)

	path := filepath.Join(ex, "testinput.txt")
	data, err := generics.ReadLines(path)
	generics.Check(err)

	textString := data[0]
	// fmt.Println(textString)

	return textString
}

func main() {
	fmt.Println("Result A:", partA())
	// fmt.Println("Result B:", partB())
}
