# Array in Golang
### declaring an array
```go
var fruitList [4]string 
```
- In this case we dont need to initialize it on the place 

### Initializing an array with values
```go
vegetableList := [3]string{"Carrot", "Potato", "Tomato"}
fmt.Println("Vegetable List:", vegetableList)
```

Output:
```
Vegetable List: [Carrot Potato Tomato]
```
**What if no value is given at some index**
```go
 var fruitList [4]string 

    // Assigning values to the array
    fruitList[0] = "Apple"
    fruitList[1] = "Orange"
    fruitList[3] = "Mango"

```
- Here `fruitList[2]` will hold an empty string because that is the zero value of the string
--------------------------------------------------------------------------------------------
### Using for loop to iterate over the array
```go
 fmt.Println("Iterating through the fruit list:")
    for i := 0; i < len(fruitList); i++ {
        fmt.Printf("fruitList[%d]: %s\n", i, fruitList[i])
    }
```
```
Iterating through the fruit list:
vegetableList[0]: Carrot
vegetableList[1]: Potato
vegetableList[2]: Tomato
```


### Iterating using the range keyword
```go
 fmt.Println("Iterating through the vegetable list using range:")
    for index, value := range vegetableList {
        fmt.Printf("vegetableList[%d]: %s\n", index, value)
    }
```

```
Iterating through the vegetable list using range:
vegetableList[0]: Carrot
vegetableList[1]: Potato
vegetableList[2]: Tomato
```

--------------------------------------------------------------------------------------------

### Copying Arrays
```go
    var copiedFruitList [4]string
    copiedFruitList = fruitList
    fmt.Println("Copied Fruit List:", copiedFruitList)

```

### Slicing an array
```go
var myArray = [4]int{1, 2, 3, 4}

// Creating a slice from the array
mySlice := myArray[1:3]

// Printing the slice
fmt.Println("Elements from 1 to 2 are :",mySlice) 
```
```
Elements from 1 to 2 are : [2 3]
```
----------------------------------------------------------------------------------------------------------


# Array vs Slices in Golang
### This is an array declaration and intialization
```go
vegetableList := [3]string{"Carrot", "Potato", "Tomato"}

```

### This is a slice declaration and intialization
```go
vegetableList := []string{"Carrot", "Potato", "Tomato"}

```
* Slices are dynamic in size 


# Slices in Golang


### Appending an element to a slice
```go
  var fruitList = []string{"Guava"}

   fmt.Println(fruitList)

   fruitList = append(fruitList, "Apple")
   fruitList = append(fruitList, "Orange")
   fruitList = append(fruitList, "Grapes")

   fruitList = append(fruitList, "Banana","strawberry") // we can append multiple elements

   fmt.Println(fruitList)
```
```
[Guava]
[Guava Apple Orange Grapes Banana strawberry]
```
                       0      1     2     3       4       5
* Original slice = `[Guava Apple Orange Grapes Banana strawberry]`

```go
fruitList = append(fruitList[2:3])
fmt.Println(fruitList)
```

```
[Orange Grapes Banana strawberry]
```
                       0      1     2     3       4       5
* Original slice = `[Guava Apple Orange Grapes Banana strawberry]`
```go
fruitList = append(fruitList[2:5]) // elements from 2nd index to 4th index
fmt.Println(fruitList)
```
```
[Orange Grapes Banana]
```