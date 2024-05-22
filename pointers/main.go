package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of poiners")

	var number int = 34234

	var ptr *int = &number

	fmt.Println(number)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	

	var ptr2 *int ;
	fmt.Println(ptr2)
	
}