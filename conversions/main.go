package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	fmt.Println("Welcome to our Pizza app")
	fmt.Println("Please rate our pizza between 1 and 5 ")

	reader:=bufio.NewReader(os.Stdin)

	input,_ :=reader.ReadString('\n')
	fmt.Println("Thanks for rating,  ",input)
}