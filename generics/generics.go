package generics

import (
	"bufio"
	"fmt"
	"os"
)

// Hello returns a greeting for the named person.
func Hello(name string) string {
	// Return a greeting that embeds the name in a message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}

func AbsInt(x int) int {
	return AbsDiffInt(x, 0)
}

func AbsDiffInt(x, y int) int {
	if x < y {
		return y - x
	}
	return x - y
}

func ReadLines(path string) ([]string, error) {
	// Reads the lines of a file as individual items in a slice and
	// returns them, then closes the file.
	file, err := os.Open(path)
	Check(err)

	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}
