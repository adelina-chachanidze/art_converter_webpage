package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func decodeArt(input string) string {
	// Regular expression to match patterns like [5 D]
	pattern := regexp.MustCompile(`\[(\d+)\s+([^\]]+)\]`)

	// Replace all matches with the expanded string
	result := pattern.ReplaceAllStringFunc(input, func(match string) string {
		// Extract the number and character from the match
		parts := pattern.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}

		// Convert the number string to integer
		count, err := strconv.Atoi(parts[1])
		if err != nil {
			return match
		}

		// Repeat the character count times
		return strings.Repeat(parts[2], count)
	})

	return result
}

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run . \"[pattern]\"")
		fmt.Println("Example: go run . \"[5 #][5 -_]-[5 #]\"")
		os.Exit(1)
	}

	input := os.Args[1]
	result := decodeArt(input)
	fmt.Println(result)
}
