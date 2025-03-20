package main

import (
	"fmt"
)

func showOperationMenu() string {
	fmt.Println("\n1. Encode")
	fmt.Println("2. Decode")
	fmt.Println("3. Exit")
	fmt.Print("Choose operation (1-3): ")
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
	fmt.Println("Welcome to the Art Encoder/Decoder!")

	for {
		choice := showOperationMenu()

		switch choice {
		case "1", "2":
			var input string
			if choice == "1" {
				fmt.Print("Enter the text to encode: ")
				fmt.Scanln(&input)
				fmt.Println("\nEncoded result:")
				fmt.Println(encodeArt(input))
			} else {
				fmt.Print("Enter the pattern to decode: ")
				fmt.Scanln(&input)
				fmt.Println("\nDecoded result:")
				fmt.Println(decodeArt(input))
			}

			// After operation, show continue/exit menu
			continueChoice := showContinueMenu()
			if continueChoice == "2" {
				fmt.Println("Thank you for using Art Encoder/Decoder. Come back soon!")
				return
			}
			// If they choose to continue, skip the welcome message
			// and just show the operation menu in the next loop

		case "3":
			fmt.Println("Thank you for using Art Encoder/Decoder. Come back soon!")
			return

		default:
			fmt.Println("Invalid choice. Please select 1, 2, or 3.")
		}
	}
}
