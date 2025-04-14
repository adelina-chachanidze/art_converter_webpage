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
	http.HandleFunc("/", handleMainPage)
	http.HandleFunc("/decoder", handleDecoder)
	http.HandleFunc("/decode-page", handleDecodePage)
	http.HandleFunc("/encode", handleEncode)
	http.HandleFunc("/decode", handleDecode)
	http.HandleFunc("/copy-encoded", handleCopyEncoded)
	http.HandleFunc("/copy-decoded", handleCopyDecoded)

	// Start the server
	fmt.Println("Server is running at http://localhost:8080")
	fmt.Println("Press Ctrl+C to stop the server...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Error starting server:", err)
		os.Exit(1)
	}
}

// handleMainPage returns the main page of the web interface with a 200 status
func handleMainPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// Return HTTP 200 status for GET requests
	if r.Method == http.MethodGet {
		w.WriteHeader(http.StatusOK)
	}

	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, PageData{})
}

// handleDecoder handles POST requests to /decoder endpoint
func handleDecoder(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	encodedString := r.FormValue("input")

	// Check for malformed encoded strings
	if err := errorsDecoding(encodedString); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Process valid encoded strings
	decoded := decodeArt(encodedString)

	// Remove only leading and trailing newlines, but preserve spaces
	decoded = strings.Trim(decoded, "\n")

	// The WriteHeader must be called before any Write operation
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusAccepted)
	w.Write([]byte(decoded))
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

	// Remove only leading and trailing newlines, but preserve spaces
	encoded = strings.Trim(encoded, "\n")

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

	// Remove only leading and trailing newlines, but preserve spaces
	decoded = strings.Trim(decoded, "\n")

	tmpl, err := template.ParseFiles("decode.html")
	if err != nil {
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, PageData{DecodedResult: decoded})
}

// Handler for copying encoded result
func handleCopyEncoded(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")

	// Set headers to tell the browser this is a text download
	w.Header().Set("Content-Disposition", "attachment; filename=encoded_art.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))

	// Write the content
	w.Write([]byte(content))
}

// Handler for copying decoded result
func handleCopyDecoded(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error parsing form", http.StatusBadRequest)
		return
	}

	content := r.FormValue("content")

	// Set headers to tell the browser this is a text download
	w.Header().Set("Content-Disposition", "attachment; filename=decoded_art.txt")
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))

	// Write the content
	w.Write([]byte(content))
}
