
## Input Methods

This program uses two different methods to capture user input:

### fmt.Scanln

**Usage:**
```go
fmt.Println("Enter your name:")
var name string
fmt.Scanln(&name)
```

**Description:**
- `fmt.Scanln` reads a single line of input from the standard input (`stdin`) and stops reading at the newline character (`\n`).
- It scans the input based on whitespace and stores the first input token in the provided variable.
- Suitable for simple, single-line inputs without spaces.

**Pros:**
- Simple and straightforward for basic inputs.
- Easy to use for reading single-line inputs without spaces.

**Cons:**
- Limited in handling complex input scenarios, such as multiple words or lines.
- Doesn't provide much control over the input processing.

### bufio.NewReader and ReadString

**Usage:**
```go
reader := bufio.NewReader(os.Stdin)
fmt.Println("Enter the rating of our Pizza:")
input, _ := reader.ReadString('\n')
input = input[:len(input)-1] // Removing the newline character
```

**Description:**
- `bufio.NewReader` creates a buffered reader that can read from an `io.Reader`, such as `os.Stdin`.
- `ReadString` reads input from the buffered reader until it encounters the specified delimiter (newline character `\n`).
- Captures the entire line of input, including spaces, making it suitable for more complex input scenarios.

**Pros:**
- Handles complex input scenarios, including multiple words and lines.
- Provides more control over the input processing and allows for more sophisticated input handling.

**Cons:**
- Slightly more complex to set up and use compared to `fmt.Scanln`.
- Requires handling of the newline character and potential errors.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
```

This README file is now ready to be used in your GitHub repository with proper code snippets and syntax highlighting.
