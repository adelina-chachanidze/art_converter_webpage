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
	fmt.Println("\n1. Encode: Convert repeated characters to shortened format")
	fmt.Println("2. Decode: Expand shortened art back to normal")
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

func main() {
	fmt.Println("Welcome to the Art Encoder/Decoder Tool!")

	for {
		choice := showOperationMenu()

		switch choice {
		case "1", "2":
			if choice == "1" {
				fmt.Println("Enter the text to encode (press Enter twice to finish):")
				input := readMultiLineInput()
				fmt.Println("\nEncoded result:")
				fmt.Println(encodeArt(input))
			} else {
				fmt.Println("Enter the pattern to decode (press Enter twice to finish):")
				input := readMultiLineInput()
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
