# Art Encoder/Decoder

A web application for encoding and decoding ASCII art. This tool helps compress ASCII art by identifying repeating patterns and representing them in a compact form.

## Features

- Encode ASCII art into a compressed format
- Decode compressed format back to ASCII art
- Supports a wide range of symbols and Unicode characters
- Recognizes and efficiently compresses complex patterns
- Save encoded/decoded results as .txt files
- Web interface with intuitive design for easy use
- Supports both one- and multi-line inputs

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

The go.mod file is not included in this repository, so you'll need to set up the Go module before running the application:


# Initialize a new Go module
```
go mod init art
```

## Running the Server

Start the server with:

```
go run main.go decoding.go
```

The server will start at http://localhost:8080. You should see the following output:

```
Server is running at http://localhost:8080
Press Ctrl+C to stop the server...
```

## Usage

1. **Encoding ASCII Art**:
   - Visit http://localhost:8080/
   - Paste or type your ASCII art in the input field
   - Click "Encode"
   - The encoded result will appear in the output field
   - Click "Save as .txt" to download the result as a text file

2. **Decoding**:
   - Visit http://localhost:8080/decoder or click "Decoder" in the navigation
   - Paste or type the encoded art in the input field
   - Input must be [number space symbol(s)], eg [8 #h]; otherwise the error will appear
   - Click "Decode"
   - The original ASCII art will appear in the output field
   - Click "Save as .txt" to download the result as a text file

## Test Examples

The repository includes a `tests` folder with several examples you can try:

- Files with `.art.txt` extension contain the original ASCII art
- Files with `.encoded.txt` extension contain the encoded version

You can use these examples to test the application's encoding and decoding capabilities.

## User Interface

The application features a clean, modern user interface with:

- Responsive design
- Intuitive navigation between encoder and decoder pages
- Large text areas for easy input and viewing of results
- Error messages that provide clear guidance on fixing invalid inputs

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
