# Different ways to use variables in Golang


**Method 1**
```go
var username string = "Abhinav"
fmt.Println(username)
fmt.Printf("Variable is of type :  %T\n", username)
```
**Default Values**
```go
var anothervariable int
fmt.Println(anothervariable) 
fmt.Printf("Variable is of type :  %T\n", anothervariable)
```
- Go does not put garbage values in the variables instead the variable anothervariable will be initialized with zero

**Method 2 (Implicit type)**
```go
var website = "This is my string"
fmt.Println(website)
fmt.Printf("Variable is of type :  %T\n", website)
```
- Here we dont need to write the datatype after the variable name 

**Method 3 (using Walrus operator)**
```go
numberOfUsers := 300000
fmt.Println(numberOfUsers)
fmt.Printf("Variable is of type :  %T\n", numberOfUsers)
```
- Here we dont need to write var before the wariable name


## Const in Golang
```go
const LoginToken string = "abcdefghijklmnopqrstuvwxyzABCDEF"
```
- Here the L of the LoginToken is capital this makes it a public variable
- In Go , a variable declared with an initial capital letter , making it public , is accesible from any other package that
import the package where the variable is defined 

