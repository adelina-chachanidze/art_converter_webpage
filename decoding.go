package main

import (
	"regexp"
	"strconv"
	"strings"
)

func decodeArt(input string) string {
	// Regular expression to match patterns like [22  ] or [4 _]
	pattern := regexp.MustCompile(`\[(\d+)([^]]+)\]`)

	// Replace all matches with the expanded string
	result := pattern.ReplaceAllStringFunc(input, func(match string) string {
		// Extract the number and pattern from the match
		parts := pattern.FindStringSubmatch(match)
		if len(parts) != 3 {
			return match
		}

		// Convert the number string to integer
		count, err := strconv.Atoi(parts[1])
		if err != nil {
			return match
		}

		// Get the pattern, removing just the first space after the number
		pattern := strings.TrimPrefix(parts[2], " ")

		// Special case: if pattern is only whitespace
		if strings.TrimSpace(pattern) == "" {
			return strings.Repeat(" ", count)
		}

		// For all other cases, just repeat the exact pattern
		return strings.Repeat(pattern, count)
	})

	return result
}

func encodeArt(input string) string {
	return "encoding is under construction"
}
