package main

import "fmt"

func main() {
	fmt.Println("Welcome to a class of poiners")

	var number int = 10

	var ptr *int = &number

	fmt.Println(number)
	fmt.Println(ptr)
	fmt.Println(*ptr)
	

	var ptr2 *int ;
	fmt.Println(ptr2)


	*ptr = *ptr * 3
	fmt.Println(number)
	
}