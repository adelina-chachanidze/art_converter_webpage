package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func decodeArt(input string) string {
	// Regular expression to match patterns like [22  ] or [4 h&]
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

// Find the best repeating pattern at the current position
func findRepeatingPattern(line string, startPos int) (pattern string, count int, endPos int) {
	maxPatternLen := 30 // Maximum pattern length to search for
	if startPos+maxPatternLen > len(line) {
		maxPatternLen = len(line) - startPos
	}

	// First check for single character repetition including Unicode characters
	if startPos < len(line) {
		// Get the rune at the current position
		r, size := utf8.DecodeRuneInString(line[startPos:])
		if r != utf8.RuneError {
			// Count how many times this rune repeats
			runeCount := 1
			j := startPos + size

			for j < len(line) {
				nextR, nextSize := utf8.DecodeRuneInString(line[j:])
				if nextR != r {
					break
				}
				runeCount++
				j += nextSize
			}

			// If we have a significant single character repetition, prefer it
			if runeCount > 1 {
				return string(r), runeCount, j
			}
		}
	}

	// Check for special paired patterns like ^|^|^|^ or | | |
	if startPos+2 <= len(line) {
		// Try to find pattern like ^|^|^|^ (two-character pattern)
		twoCharPattern := line[startPos : startPos+2]
		pairCount := 1
		j := startPos + 2

		for j+2 <= len(line) && line[j:j+2] == twoCharPattern {
			pairCount++
			j += 2
		}

		if pairCount > 1 {
			return twoCharPattern, pairCount, j
		}
	}

	bestPattern := ""
	bestCount := 0
	bestEndPos := startPos
	bestSavings := 0

	// Try patterns of different lengths
	for patternLen := 2; patternLen <= maxPatternLen; patternLen++ {
		if startPos+patternLen > len(line) {
			continue
		}

		pattern := line[startPos : startPos+patternLen]

		// Skip patterns that are just spaces
		if strings.TrimSpace(pattern) == "" {
			continue
		}

		// Count how many times this pattern repeats
		count := 1
		pos := startPos + patternLen

		for pos+patternLen <= len(line) && line[pos:pos+patternLen] == pattern {
			count++
			pos += patternLen
		}

		// If pattern repeats more than once, calculate savings
		if count >= 2 {
			// Calculate space saved by using [n pattern] format
			// Original: patternLen * count
			// Encoded: 4 + len(count) + patternLen (for [n pattern])
			originalSize := patternLen * count
			encodedSize := 4 + len(strconv.Itoa(count)) + patternLen
			savings := originalSize - encodedSize

			// Check if this is the best pattern so far
			if savings > bestSavings || (savings == bestSavings && patternLen < len(bestPattern)) {
				bestPattern = pattern
				bestCount = count
				bestEndPos = pos
				bestSavings = savings
			}
		}
	}

	// If we found a good pattern, return it
	if bestSavings > 0 {
		return bestPattern, bestCount, bestEndPos
	}

	return "", 0, startPos
}

func encodeArt(input string) string {
	// Only trim newlines, preserve spaces
	input = strings.Trim(input, "\n")
	lines := strings.Split(input, "\n")
	var result []string

	for _, line := range lines {
		var encoded strings.Builder
		i := 0

		// Special case for lines with all the same character (like "######")
		if len(line) > 0 {
			// Check for a homogeneous line (all same character)
			r, size := utf8.DecodeRuneInString(line)
			if r != utf8.RuneError {
				allSame := true
				j := size

				for j < len(line) {
					nextR, nextSize := utf8.DecodeRuneInString(line[j:])
					if nextR != r {
						allSame = false
						break
					}
					j += nextSize
				}

				if allSame {
					count := utf8.RuneCountInString(line)
					encoded.WriteString(fmt.Sprintf("[%d %s]", count, string(r)))
					result = append(result, encoded.String())
					continue
				}
			}
		}

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
			// Look for repeating patterns
			pattern, count, newPos := findRepeatingPattern(line, i)
			if count > 0 {
				encoded.WriteString(fmt.Sprintf("[%d %s]", count, pattern))
				i = newPos
				continue
			}

			// No repeating pattern found, just output the current character
			// Ensure we correctly handle Unicode characters by getting rune
			r, size := utf8.DecodeRuneInString(line[i:])
			encoded.WriteString(string(r))
			i += size
		}

		result = append(result, encoded.String())
	}

	return strings.Join(result, "\n")
}
