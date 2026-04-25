# Go CLI Utilities

A collection of command-line programs built in Go that practice core concepts like file processing, text transformation, input validation, and unit testing.

---

## Programs

### 1. Validate ASCII Input

Checks if every character in a string is a printable ASCII character (between `' '` and `'~'`).

**Usage**
```bash
go run . "Hello!"
```

**Output**
```
Valid input
```

**Invalid input**
```bash
go run . "Hello😊"
```
```
Invalid character found: 😊
```

**Concepts Practiced**
- `rune` and ASCII range checking
- `for range` loop over strings
- Input validation
- `os.Args`

---

### 2. Replace a Word in a File

Reads a file and replaces every occurrence of one word with another, then prints the result.

**Usage**
```bash
go run . input.txt oldWord newWord
```

**Example**

`input.txt`:
```
I love Java
Java is popular
```

```bash
go run . input.txt Java Go
```

**Output**
```
I love Go
Go is popular
```

**Concepts Practiced**
- `os.Args`
- `strings.ReplaceAll`
- File processing with `os.ReadFile`
- Input validation
- Error handling

---

### 3. Count Words in a File

Reads a file and counts the total number of words inside it.

**Usage**
```bash
go run . <filename>
```

**Example**

`input.txt`:
```
Go is simple and powerful
```

**Output**
```
Total words: 5
```

**Concepts Practiced**
- `strings.Fields`
- `len`
- File processing
- `os.Args`

---

### 4. Split Text by Newline

Receives a string from the terminal and splits it by `\n`, printing each segment on its own line.

**Usage**
```bash
go run . "Hello\nThere"
```

**Output**
```
Line 1: Hello
Line 2: There
```

**Concepts Practiced**
- `os.Args`
- `strings.Split`
- Escaped newline handling (`\\n`)

---

### 5. Count Words (with Unit Tests)

A standalone function that counts words in a string, with full table-driven unit tests.

**Function signature**
```go
func CountWords(text string) int
```

**Examples**
| Input | Output |
|---|---|
| `"Go is powerful"` | `3` |
| `""` | `0` |
| `"Hello   there"` | `2` |
| `"One\nTwo\nThree"` | `3` |

**Run tests**
```bash
go test -v ./...
```

**Concepts Practiced**
- `testing` package
- `go test`
- Table-driven tests
- `strings.Fields`

## Project Structure

```
.
├── main.go          # entry point and all functions
├── main_test.go     # table-driven unit tests
├── input.txt        # sample input file
└── README.md
```

---

## Running the Programs

```bash
# run with input validation
go run . "Hello World"

# replace word in file
go run . input.txt Java Go

# run all tests
go test ./...

# run tests with verbose output
go test -v ./...

# run a specific test
go test -run TestCountWords/empty_string -v
```

---

## Functions Reference

| Function | Description |
|---|---|
| `ValidateASCIIInput(str string)` | Checks every character is a printable ASCII character |
| `ReplaceWordInFile(file, newWord, oldWord string) string` | Replaces a word in a file and returns the result |
| `ReplaceAllInText(text, oldWord, newWord string) string` | Replaces all occurrences of a word in a string |
| `CountWords(text string) int` | Returns the number of words in a string |
| `CountLinesInText(input string) int` | Returns the number of lines in a string |
| `PrintLineSegments(str string) error` | Splits a string by `\n` and prints each segment |