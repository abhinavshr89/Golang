package main

import "fmt"

func main() {
    fmt.Println("Hello World")

    // Declaring an array of strings with length 4
    var fruitList [4]string 

    // Assigning values to the array
    fruitList[0] = "Apple"
    fruitList[1] = "Orange"
    fruitList[3] = "Mango"

    // Printing the entire array
    fmt.Println("Fruit List:", fruitList)

    // Printing a specific element (this will be an empty string)
    fmt.Println("Element at index 2:", fruitList[2])

    // Initializing an array with values
    vegetableList := [3]string{"Carrot", "Potato", "Tomato"}
    fmt.Println("Vegetable List:", vegetableList)

    // Iterating through the array using a for loop
    fmt.Println("Iterating through the fruit list:")
    for i := 0; i < len(vegetableList); i++ {
        fmt.Printf("vegetableList[%d]: %s\n", i, vegetableList[i])
    }

    // // Iterating using the range keyword
    fmt.Println("Iterating through the vegetable list using range:")
    for index, value := range vegetableList {
        fmt.Printf("vegetableList[%d]: %s\n", index, value)
    }

    // // Copying arrays
    var copiedFruitList [4]string
    copiedFruitList = fruitList
    fmt.Println("Copied Fruit List:", copiedFruitList)

    // // Slicing an array (creating a slice from an array)
	var myArray = [4]int{1, 2, 3, 4}

    // Creating a slice from the array
    mySlice := myArray[1:3]

    // Printing the slice
    fmt.Println("Elements from 1 to 2 are :",mySlice) 

}
