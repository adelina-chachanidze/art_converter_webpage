package main

import (
	//"bufio"
	//"fmt"
	//"os"
	"regexp"
	"strconv"
	"strings"
)

func processInput(input string) string {
	lines := strings.Split(input, "\n")
	var outputLines []string

	for _, line := range lines {
		processedLine := processLine(line)
		outputLines = append(outputLines, processedLine)
	}

	return strings.Join(outputLines, "\n")
}

func processLine(line string) string {
	var result strings.Builder
	current := 0

	for current < len(line) {
		// Look for patterns like [19  ]__
		spaceCountMatch := regexp.MustCompile(`\[(\d+)\s+\]`).FindStringSubmatchIndex(line[current:])
		if spaceCountMatch != nil && spaceCountMatch[0] == 0 {
			countStr := line[current+spaceCountMatch[2]:current+spaceCountMatch[3]]
			count, _ := strconv.Atoi(countStr)
			result.WriteString(strings.Repeat(" ", count))
			current += spaceCountMatch[1]
			continue
		}

		// Look for patterns like [9 o]
		charCountMatch := regexp.MustCompile(`\[(\d+)\s+([^]]+)\]`).FindStringSubmatchIndex(line[current:])
		if charCountMatch != nil && charCountMatch[0] == 0 {
			countStr := line[current+charCountMatch[2]:current+charCountMatch[3]]
			char := line[current+charCountMatch[4]:current+charCountMatch[5]]
			count, _ := strconv.Atoi(countStr)
			result.WriteString(strings.Repeat(char, count))
			current += charCountMatch[1]
			continue
		}

		// For any other characters (like *, "", etc.)
		if current < len(line) {
			result.WriteByte(line[current])
			current++
		}
	}

	// Post-process any special patterns
	processed := result.String()
	
	// Process double quotes
	processed = strings.ReplaceAll(processed, "\"\"", "\"\"")
	
	return processed
}