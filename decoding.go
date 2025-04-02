package main

import (
	"fmt"
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
	// Split input into lines
	lines := strings.Split(input, "\n")
	var result []string

	for _, line := range lines {
		if line == "" {
			continue
		}

		var encoded strings.Builder
		count := 1
		currentChar := line[0]

		// Handle the case where line is all the same character
		allSame := true
		for i := 1; i < len(line); i++ {
			if line[i] != currentChar {
				allSame = false
				break
			}
		}
		if allSame && len(line) > 0 {
			result = append(result, fmt.Sprintf("[%d %c]", len(line), currentChar))
			continue
		}

		// Process character by character
		for i := 1; i <= len(line); i++ {
			// If we're at the end or found a different character
			if i == len(line) || line[i] != currentChar {
				// If count is greater than 3, use compression
				if count > 3 {
					encoded.WriteString(fmt.Sprintf("[%d %c]", count, currentChar))
				} else {
					// Otherwise write characters directly
					for j := 0; j < count; j++ {
						encoded.WriteByte(currentChar)
					}
				}

				if i < len(line) {
					count = 1
					currentChar = line[i]
				}
			} else {
				count++
			}
		}

		result = append(result, encoded.String())
	}

	return strings.Join(result, "\n")
}
