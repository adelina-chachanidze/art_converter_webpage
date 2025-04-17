# Art Encoder/Decoder

A web application for encoding and decoding ASCII art. This tool helps compress ASCII art by identifying repeating patterns and representing them in a compact form.

## Features

- Encode ASCII art into a compressed format
- Decode compressed format back to ASCII art
- Supports a wide range of symbols and Unicode characters
- Recognizes and efficiently compresses complex patterns
- Save encoded/decoded results as .txt files
- Web interface with intuitive design for easy use
- Copy encoded/decoded results with a single click

## Requirements

- Go 1.23.2 or higher

## Installation

1. Clone this repository:
   ```
   git clone https://gitea.koodsisu.fi/adelinachachanidze/art.git
   cd art
   ```

2. Ensure you have Go installed:
   ```
   go version
   ```
   The output should show Go version 1.23.2 or higher.

## Building the Project

No build step is required as this is a Go application that can be run directly.

The go.mod file is not included in this repository, so you'll need to set up the Go module before running the application:

```
# Initialize a new Go module
go mod init art

## Running the Server

Start the server with:

```
go run .
```

The server will start at http://localhost:8080. You should see the following output:

```
Server is running at http://localhost:8080
Press Ctrl+C to stop the server...
```

To end the session press Ctrl+C

## Usage

1. **Encoding ASCII Art**:
   - Visit http://localhost:8080/
   - Paste your ASCII art in the input field
   - Click "Encode"
   - The encoded result will appear in the output field
   - Use the "Copy" button to copy the encoded result
   - Click "Save as TXT" to download the result as a text file

2. **Decoding**:
   - Visit http://localhost:8080/decoder or click "Decoder" in the navigation
   - Paste the encoded art in the input field
   - Click "Decode"
   - The original ASCII art will appear in the output field
   - Use the "Copy" button to copy the decoded result
   - Click "Save as TXT" to download the result as a text file

## User Interface

The application features a clean, modern user interface with:

- Responsive design that works on both desktop and mobile devices
- Intuitive navigation between encoder and decoder pages
- Large text areas for easy input and viewing of results
- Error messages that provide clear guidance on fixing invalid inputs
- Syntax highlighting for encoded patterns
- Light/dark mode toggle to reduce eye strain during extended use
- Real-time feedback as you interact with the application

## Encoding Format

The encoding format uses brackets to represent repeated patterns:
- `[n c]` where `n` is the number of repetitions and `c` is the character or pattern to repeat
- Example: `[5 #]` expands to `#####`
- Example: `[3 ^|]` expands to `^|^|^|`
- Example: `[4 ♥♦]` expands to `♥♦♥♦♥♦♥♦`

The encoder is smart enough to detect:
- Simple character repetition (like `-----`)
- Pattern repetition (like `^|^|^|^|`)
- Mixed symbol sequences with optimal compression
- Unicode characters and special symbols
- Space patterns with proper preservation

## Technical Details

Built with:

- Go backend for processing ASCII art
- HTML/CSS for the user interface
- Server-side rendering with Go templates
- Pure Go implementation without JavaScript dependencies

## Credits

- Background image by [Maxim Berg](https://unsplash.com/@maxberg) on [Unsplash](https://unsplash.com/photos/a-blurry-image-of-a-multicolored-background-PiFzbqDClGk)

## License

This project is licensed under the MIT License - see the LICENSE file for details.