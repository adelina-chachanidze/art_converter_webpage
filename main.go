// This file is renamed to cli.go to separate CLI and server functionalities.

package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
	"strings"
)

// Data structure for passing to templates
type PageData struct {
	EncodedResult string
	DecodedResult string
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
	// Create a file server to serve static files
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/styles.css", fs)

	// Define route handlers
	http.HandleFunc("/", handleEncodePage)
	http.HandleFunc("/decode-page", handleDecodePage)
	http.HandleFunc("/encode", handleEncode)
	http.HandleFunc("/decode", handleDecode)

	// Start the server
	fmt.Println("Server is running at http://localhost:8080")
	fmt.Println("Press Ctrl+C to stop the server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}

func handleEncodePage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, PageData{})
}

func handleDecodePage(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("decode.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, PageData{})
}

func handleEncode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	input := r.FormValue("input")

	if err := errorsEncoding(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	encoded := encodeArt(input)

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, PageData{EncodedResult: encoded})
}

func handleDecode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	input := r.FormValue("input")

	if err := errorsDecoding(input); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	decoded := decodeArt(input)

	tmpl, err := template.ParseFiles("decode.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, PageData{DecodedResult: decoded})
}
