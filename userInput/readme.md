
Feel free to copy and paste this Markdown content into your README file on GitHub. Don't forget to replace `yourusername` with your actual GitHub username and update the project URL accordingly.
```

The above Markdown example uses the following:

- **Triple backticks** (\`\`\`) to start and end the code block.
- **Language specifier** (e.g., `go`, `sh`) immediately after the opening backticks to enable syntax highlighting.

### Example for Go Code

To include a Go code snippet with syntax highlighting:

\```go
fmt.Println("Enter your name:")
var name string
fmt.Scanln(&name)
\```

### Example for Shell Commands

To include shell commands with syntax highlighting:

\```sh
git clone https://github.com/yourusername/user-input-program.git
cd user-input-program
go run main.go
\```

Just replace the backslashes (\) with proper backticks when you use it in your actual README file. Here’s the full README again with the fenced code blocks in place:

```markdown
# User Input Program in Go

This Go program demonstrates how to handle user input using two different methods: `fmt.Scanln` and `bufio.NewReader` with `ReadString`. The program greets the user, asks for their name, and then asks for a rating of a pizza, displaying the entered rating.

## Table of Contents

- [Introduction](#introduction)
- [Prerequisites](#prerequisites)
- [Installation](#installation)
- [Usage](#usage)
- [Input Methods](#input-methods)
  - [fmt.Scanln](#fmt-scanln)
  - [bufio.NewReader and ReadString](#bufio-newreader-and-readstring)
- [License](#license)

## Introduction

This program is a simple Go application that demonstrates how to read user input using two different methods. It serves as an educational example for beginners to understand the differences between these input handling techniques in Go.

## Prerequisites

- Go installed on your system. You can download it from the [official Go website](https://golang.org/dl/).

## Installation

1. Clone the repository:
   ```sh
   git clone https://github.com/yourusername/user-input-program.git
   ```
2. Navigate to the project directory:
   ```sh
   cd user-input-program
   ```

## Usage

Run the program using the following command:
```sh
go run main.go
```

The program will:
1. Display a welcome message.
2. Ask for your name.
3. Greet you by name.
4. Ask for a rating of the pizza.
5. Display the entered rating.

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
