package main

import (
    "bufio"
    "fmt"
    "os"
)

func main() {
    welcome := "Welcome to user input program"
    fmt.Println(welcome)

	//1. First Method 
    fmt.Println("Enter your name:")
    var name string
    fmt.Scanln(&name)
    fmt.Println("Hello", name)


	//2. Second Method

    reader := bufio.NewReader(os.Stdin)
    fmt.Println("Enter the rating of our Pizza:")

    // comma ok || err ok syntax
    input, err := reader.ReadString('\n')
    if err != nil {
        fmt.Println("An error occurred while reading input:", err)
        return
    }
    // Trim the newline character from the input
    input = input[:len(input)-1]
    fmt.Println("You rated the Pizza:", input)
}
