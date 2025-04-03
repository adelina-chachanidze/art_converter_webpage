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
	lines := strings.Split(input, "\n")
	var result []string

	for _, line := range lines {
		if line == "" {
			continue
		}

		var encoded strings.Builder
		i := 0

		// Count leading spaces
		spaceCount := 0
		for i < len(line) && line[i] == ' ' {
			spaceCount++
			i++
		}
		if spaceCount > 0 {
			encoded.WriteString(fmt.Sprintf("[%d  ]", spaceCount))
		}

		// Process the rest of the line
		for i < len(line) {
			// Check for repeating patterns first
			if i+1 < len(line) && (line[i] == '|' || line[i] == '^') {
				pattern := ""
				count := 0

				// Try to find repeating pattern
				if i+1 < len(line) {
					pattern = line[i : i+2]
					for j := i; j < len(line)-1; j += 2 {
						if j+1 >= len(line) || line[j:j+2] != pattern {
							break
						}
						count++
					}
				}

				if count > 2 {
					encoded.WriteString(fmt.Sprintf("[%d %s]", count, pattern))
					i += count * 2
					continue
				}
			}

			// Handle regular character sequences
			char := line[i]
			count := 1
			j := i + 1
			for j < len(line) && line[j] == char {
				count++
				j++
			}

			if count > 3 || (char == ' ' && count > 1) || char == '#' {
				encoded.WriteString(fmt.Sprintf("[%d %c]", count, char))
			} else {
				for k := 0; k < count; k++ {
					encoded.WriteByte(char)
				}
			}
			i = j
		}

		result = append(result, encoded.String())
	}

	return strings.Join(result, "\n")
}
