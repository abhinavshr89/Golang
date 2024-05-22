package main

import "fmt"

// jwtToken := 8083405  ! Wrong Way

var jwtToken int = 8083405

const LoginToken string = "abcdefghijklmnopqrstuvwxyzABCDEF"

/*Note --> Here the L of the LoginToken is capital this makes it a public variable
In Go , a variable declared with an initial capital letter , making it public , is accesible from any other package that
import the package where the variable is defined */

func main() {
	var username string = "Abhinav"
	fmt.Println(username)
	fmt.Printf("Variable is of type :  %T\n", username)

	var isLoggedIn bool = false
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type :  %T\n", isLoggedIn)

	var smallVal int = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type :  %T\n", smallVal)

	var smallFloat float64 = 234.32432423543434324
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type :  %T\n", smallFloat)

	//default values and some aliases

	var anothervariable int
	fmt.Println(anothervariable) // --> Output : 0 (Go does not put garbage values in the variables)
	fmt.Printf("Variable is of type :  %T\n", anothervariable)

	//implicit type

	var website = "This is my string"
	fmt.Println(website)
	fmt.Printf("Variable is of type :  %T\n", website)

	//no var style

	numberOfUsers := 300000
	fmt.Println(numberOfUsers)
	fmt.Printf("Variable is of type :  %T\n", numberOfUsers)

	//Note --> We are allowed to use the walrus operator inside a method only we can not use this outside a method

	fmt.Println(jwtToken)
	fmt.Println(LoginToken)
}
