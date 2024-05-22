package main

import "fmt"

func main() {

   var fruitList = []string{"Guava"}
   
   fmt.Println(fruitList)

   fruitList = append(fruitList, "Apple")
   fruitList = append(fruitList, "Orange")
   fruitList = append(fruitList, "Grapes")

   fruitList = append(fruitList, "Banana","strawberry") // we can append multiple elements

   fmt.Println(fruitList)

   fruitList = append(fruitList[2:])
   fmt.Println(fruitList)
}