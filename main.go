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
	ErrorMessage  string
}

func errorsEncoding(input string) error {
	if strings.TrimSpace(input) == "" {
		return fmt.Errorf("Error: Input is empty, please provide some text to encode")
	}
	return nil
}

func errorsDecoding(input string) error {
	if strings.TrimSpace(input) == "" {
		return fmt.Errorf("Error: Input is empty, please provide some text to decode")
	}

	var openBrackets int
	var bracketContent strings.Builder

	for _, char := range input {
		if char == '[' {
			openBrackets++
			bracketContent.Reset()
		} else if char == ']' {
			if openBrackets == 0 {
				return fmt.Errorf("Error: Found ']' without matching '['")
			}
			openBrackets--

			// Validate content between brackets
			content := bracketContent.String()

			// Find the first space
			spaceIndex := strings.Index(content, " ")
			if spaceIndex == -1 {
				return fmt.Errorf("Error: Invalid format in brackets, must be [number space symbol(s)], eg [8 #]")
			}

			// Check if first part is a number
			number := content[:spaceIndex]
			if number == "" {
				return fmt.Errorf("Error: First part in brackets must be a number, eg [8 #]")
			}
			for _, digit := range number {
				if digit < '0' || digit > '9' {
					return fmt.Errorf("Error: First part in brackets must be a number, eg [8 #]")
				}
			}

			// Everything after the first space is the third argument
			// If it's empty or just spaces, that's still valid as spaces are the third argument
			if spaceIndex == len(content)-1 {
				return fmt.Errorf("Error: Missing third argument after space in brackets")
			}

		} else if openBrackets > 0 {
			bracketContent.WriteRune(char)
		}
	}

	if openBrackets > 0 {
		return fmt.Errorf("Error: Found '[' without matching ']'")
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

// handleDecoder handles GET and POST requests to /decoder endpoint
func handleDecoder(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		tmpl, err := template.ParseFiles("decode.html")
		if err != nil {
			http.Error(w, "Error loading template", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusOK)
		tmpl.Execute(w, PageData{})
		return
	}

	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		// Render the template with the error
		tmpl, _ := template.ParseFiles("decode.html")
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, PageData{ErrorMessage: "Error parsing form data"})
		return
	}

	encodedString := r.FormValue("input")

	// Check for malformed encoded strings
	if err := errorsDecoding(encodedString); err != nil {
		// Render the template with the error
		tmpl, _ := template.ParseFiles("decode.html")
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, PageData{ErrorMessage: err.Error()})
		return
	}

	// Process valid encoded strings
	decoded := decodeArt(encodedString)

	// Remove only leading and trailing newlines, but preserve spaces
	decoded = strings.Trim(decoded, "\n")

	// Render the template with the decoded result
	tmpl, err := template.ParseFiles("decode.html")
	if err != nil {
		// This is a server error, so we'll still use http.Error for this
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	tmpl.Execute(w, PageData{DecodedResult: decoded})
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

	w.WriteHeader(http.StatusOK)
	tmpl.Execute(w, PageData{})
}

func handleEncode(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseForm()
	if err != nil {
		// Render the template with the error
		tmpl, _ := template.ParseFiles("index.html")
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, PageData{ErrorMessage: "Error parsing form data"})
		return
	}

	input := r.FormValue("input")

	if err := errorsEncoding(input); err != nil {
		// Render the template with the error
		tmpl, _ := template.ParseFiles("index.html")
		w.WriteHeader(http.StatusBadRequest)
		tmpl.Execute(w, PageData{ErrorMessage: err.Error()})
		return
	}

	encoded := encodeArt(input)

	// Remove only leading and trailing newlines, but preserve spaces
	encoded = strings.Trim(encoded, "\n")

	// Render the template with the encoded result
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		// This is a server error, so we'll still use http.Error for this
		http.Error(w, "Error loading template", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusAccepted)
	tmpl.Execute(w, PageData{EncodedResult: encoded})
}

func handleDecode(w http.ResponseWriter, r *http.Request) {
	// Redirect to the decoder endpoint
	http.Redirect(w, r, "/decoder", http.StatusSeeOther)
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
