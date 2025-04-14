# ASCII Art Encoder/Decoder

A powerful tool that helps you create and manipulate text-based art through a special compression notation. This application provides both encoding (compression) and decoding (expansion) functionality for ASCII art with a clean, modern web interface.

## Features

- **Encode Mode**: Converts repetitive ASCII art patterns into a compressed format
- **Decode Mode**: Expands compressed patterns back into full ASCII art
- **Modern Web Interface**: Clean, responsive design with intuitive controls
- **Toggle Switch**: Easy switching between encode and decode modes
- **Copy to File**: One-click option to save output as a text file
- **Multi-line Support**: Process multiple lines of art at once
- **Error Handling**: Robust error checking for invalid patterns

## Installation

Ensure you have Go installed on your system, then:

```bash
# Clone the repository
git clone https://gitea.koodsisu.fi/adelinachachanidze/art.git
cd art

# Run the web server
go run .
```

Then open your browser and navigate to: http://localhost:8080

## Usage

The web interface offers an intuitive way to work with ASCII art:

1. **Toggle between Encode/Decode**: Use the toggle switch at the top to choose mode
2. **Input**: Enter your ASCII art or compressed notation in the text area
3. **Process**: Click the "Encode" or "Decode" button depending on your mode
4. **Output**: View the result in the output area below
5. **Copy to File**: Click "Copy" in the top-right corner of the output to save as a text file

### Encoding Format

When encoding, the tool automatically detects repeating patterns and converts them into a compressed format:

- Repeated characters are converted to `[N X]` format where:
  - `N` is the number of repetitions
  - `X` is the character or pattern to repeat
- Leading spaces are automatically compressed
- Special patterns (like `|` or `^`) are detected and encoded
- Sequences of 3 or more identical characters are compressed

### Decoding Format

The tool understands the following notation:

- `[N X]`: Repeat character/pattern X exactly N times
- Characters outside brackets are printed as-is
- Spaces can be encoded as `[N  ]` (two spaces after N)

## Examples

### Basic Example
```bash
# Input (Decode Mode):
[5 #][5 -_]-[5 #]

# Output:
#####-_-_-_-_-_-#####
```

### Complex Example
```bash
# Input (Encode Mode):
AAAAAAA___BBBBB

# Output:
[7 A][3 _][5 B]
```

### Multi-line Art Example
```bash
# Input (Decode Mode):
[10  ]___
[10  ]\\ \
[10  ]\\ `\
[5  ]____[5  ]\\  \

# Output:
          ___
          \\ \
          \\ `\
     ____     \\  \
```

## Error Handling

The tool includes validation for:
- Empty input
- Unmatched brackets
- Invalid number format
- Missing arguments in brackets
- Proper bracket syntax

## Web Interface

The web interface features:
- Clean, modern design with a responsive layout
- Intuitive toggle switch to change between modes
- Syntax highlighting for input and output
- One-click copy functionality to save results as text files
- Clear error notifications for invalid input

## Technical Details

Built with:
- Go backend for processing ASCII art
- HTML/CSS for the user interface
- Server-side rendering with Go templates
- Pure Go implementation without JavaScript dependencies

## Credits

- Background image by [Maxim Berg](https://unsplash.com/@maxberg) on [Unsplash](https://unsplash.com/photos/a-blurry-image-of-a-multicolored-background-PiFzbqDClGk)