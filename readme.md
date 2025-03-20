# Art Decoder

A command-line tool that converts a string with special notation into text-based art.

## Usage

```bash
go run . "[pattern]"
```

## Pattern Syntax

The tool uses a special notation to describe consecutive characters:

- `[N X]` means repeat character X exactly N times
- Any characters outside of the brackets are printed as-is

## Examples

```bash
# Basic example
go run . "[5 #][5 -_]-[5 #]"
# Output: #####-_-_-_-_-_-#####

# More complex example
go run . "[3 @][2 -][4 #]"
# Output: @@@--####
```

## How it Works

The tool uses regular expressions to find patterns like `[N X]` and replaces them with the character X repeated N times. This makes it easy to create large text-based art pieces without having to type the same character multiple times.
