package main

import (
	"fmt"
	"sort"
)

func main() {

   var fruitList = []string{"Guava"}
   
   fmt.Println(fruitList)

   fruitList = append(fruitList, "Apple")
   fruitList = append(fruitList, "Orange")
   fruitList = append(fruitList, "Grapes")

   fruitList = append(fruitList, "Banana","strawberry") // we can append multiple elements

   fmt.Println(fruitList)

   fruitList = append(fruitList[:5])
   fmt.Println(fruitList)


   highscore := make([]int , 4)

   highscore[0] = 100
   highscore[1] = 200
   highscore[2] = 400
   highscore[3] = 250
// highscore[4]=450  !This will cause error
   fmt.Println(highscore)

   highscore = append(highscore, 50, 900, 740, 800)
   fmt.Println(highscore)
   
   sort.Ints(highscore)
   fmt.Println(highscore)
   fmt.Println(sort.IntsAreSorted(highscore))

   courses := []string{"reactjs","javascript","swift","ruby"}
   fmt.Println(courses)
   
   var index int = 2 
   courses = append(courses[:index],courses[index+1:]...)
   fmt.Println(courses)

}