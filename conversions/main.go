package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our pizza between 1 and 5 ")

	reader := bufio.NewReader(os.Stdin)

	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// Trim any leading or trailing whitespace from the input
	input = strings.TrimSpace(input)

	// Convert the input to a float64
	numRating, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Please enter a valid number between 1 and 5")
		return
	}

	// Check if the rating is within the valid range
	if numRating < 1 || numRating > 5 {
		fmt.Println("Please rate our pizza between 1 and 5")
		return
	}

	fmt.Println("Thanks for rating: ", numRating)
	fmt.Println("Added 1 to the rating: ", numRating+1)
}
