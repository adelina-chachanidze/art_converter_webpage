# Art Encoder/Decoder

A powerful command-line tool that helps you create and manipulate text-based art through a special compression notation. This tool provides both encoding (compression) and decoding (expansion) functionality for ASCII art.

## Features

- **Encode Mode**: Converts repetitive ASCII art patterns into a compressed format
- **Decode Mode**: Expands compressed patterns back into full ASCII art
- **Interactive CLI**: User-friendly menu-driven interface
- **Multi-line Support**: Process multiple lines of art at once
- **Error Handling**: Robust error checking for invalid patterns

## Installation

Ensure you have Go installed on your system, then:

```bash
# Clone the repository
git clone https://gitea.koodsisu.fi/adelinachachanidze/art.git
cd art-encoder-decoder

# Run the program
go run .
```

## Usage

The program provides an interactive menu with the following options:

1. **Encode**: Convert ASCII art to compressed format
2. **Decode**: Expand compressed format back to ASCII art
3. **Exit**: Close the program

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

## How It Works

The tool uses regular expressions to:
1. For decoding: Find patterns like `[N X]` and replace them with the expanded repetition
2. For encoding: Detect repeating patterns and convert them into the compressed notation
3. Handle special cases like leading spaces and multi-character patterns