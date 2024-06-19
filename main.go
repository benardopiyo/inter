package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args
	if len(args) != 3 {
		return // If number of arguments is not 3, do nothing
	}

	first := args[1]
	second := args[2]

	// Create a map to track characters from the first string
	firstMap := make(map[rune]bool)
	for _, char := range first {
		firstMap[char] = true
	}

	// Create variables to accumulate result and track seen characters
	var result []rune
	seen := make(map[rune]bool)

	// Iterate through characters of the second string
	for _, char := range second {
		if firstMap[char] && !seen[char] {
			result = append(result, char)
			seen[char] = true
		}
	}

	// Print the result followed by a newline
	for _, char := range result {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
