package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func readMultiLineInput() string {
	var inputLines []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		// Empty line signals end of input
		if line == "" {
			break
		}
		inputLines = append(inputLines, line)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", err)
		return ""
	}

	return strings.Join(inputLines, "\n")
}

func showOperationMenu() string {
	fmt.Println("\n1. Encode: Convert art to compressed format")
	fmt.Println("2. Decode: Expand your compressed art back to normal")
	fmt.Println("3. Exit")
	fmt.Print("\nChoose operation (1-3): ")
	var choice string
	fmt.Scanln(&choice)
	return choice
}

func showContinueMenu() string {
	fmt.Println("\n1. Continue")
	fmt.Println("2. Exit")
	fmt.Print("Choose option (1-2): ")
	var choice string
	fmt.Scanln(&choice)
	return choice
}

func errorsEncoding(input string) error {
	if strings.TrimSpace(input) == "" {
		return fmt.Errorf("\033[31merror: input is empty, please provide some text to encode\033[0m")
	}
	return nil
}

func errorsDecoding(input string) error {
	if strings.TrimSpace(input) == "" {
		return fmt.Errorf("\033[31merror: input is empty, please provide some text to decode\033[0m")
	}

	var openBrackets int
	var bracketContent strings.Builder

	for _, char := range input {
		if char == '[' {
			openBrackets++
			bracketContent.Reset()
		} else if char == ']' {
			if openBrackets == 0 {
				return fmt.Errorf("\033[31merror: found ']' without matching '['\033[0m")
			}
			openBrackets--

			// Validate content between brackets
			content := bracketContent.String()

			// Find the first space
			spaceIndex := strings.Index(content, " ")
			if spaceIndex == -1 {
				return fmt.Errorf("\033[31merror: invalid format in brackets, must be 'number space symbol(s)'\033[0m")
			}

			// Check if first part is a number
			number := content[:spaceIndex]
			if number == "" {
				return fmt.Errorf("\033[31merror: first part in brackets must be a number\033[0m")
			}
			for _, digit := range number {
				if digit < '0' || digit > '9' {
					return fmt.Errorf("\033[31merror: first part in brackets must be a number\033[0m")
				}
			}

			// Everything after the first space is the third argument
			// If it's empty or just spaces, that's still valid as spaces are the third argument
			if spaceIndex == len(content)-1 {
				return fmt.Errorf("\033[31merror: missing third argument after space in brackets\033[0m")
			}

		} else if openBrackets > 0 {
			bracketContent.WriteRune(char)
		}
	}

	if openBrackets > 0 {
		return fmt.Errorf("\033[31merror: found '[' without matching ']'\033[0m")
	}

	return nil
}

func main() {
	fmt.Println("Welcome to the Art Encoder/Decoder Tool!")

	for {
		choice := showOperationMenu()

		switch choice {
		case "1", "2":
			if choice == "1" {
				var input string
				for {
					fmt.Println("Enter the text to encode (press Enter twice to finish):")
					input = readMultiLineInput()

					if err := errorsEncoding(input); err != nil {
						fmt.Println(err)
						continue
					}
					break
				}

				encoded := encodeArt(input)
				fmt.Println("\nEncoded result:")
				fmt.Println(encoded)
			} else {
				var input string
				for {
					fmt.Println("Enter the pattern to decode (press Enter twice to finish):")
					input = readMultiLineInput()

					if err := errorsDecoding(input); err != nil {
						fmt.Println(err)
						continue
					}
					break
				}

				fmt.Println("\nDecoded result:")
				fmt.Println(decodeArt(input))
			}

			continueChoice := showContinueMenu()
			if continueChoice == "2" {
				fmt.Println("Thank you for using Art Encoder/Decoder. Come back soon!")
				return
			}

		case "3":
			fmt.Println("Thank you for using Art Encoder/Decoder. Come back soon!")
			return

		default:
			fmt.Println("Invalid choice. Please select 1, 2, or 3.")
		}
	}
}
