package main

import "fmt"

const LoginToken = "12345" // publically accessible coz start with Capital(LoginToken)

func main() {
	var roleNum int = 121
	fmt.Println(roleNum)
	roleNum := 121


	var username string = "Akash"
	fmt.Println("Hello, ", username)
	fmt.Printf("variable is of type: %T \n ", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("variable is of type: %T \n ", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("variable is of type: %T \n ", smallVal)

	var smallFloat float64 = 255.4554451222
	fmt.Println(smallFloat)
	fmt.Printf("variable is of type: %T \n ", smallFloat)

	//default values and some aliases
	var anatherVariable int
	fmt.Println(anatherVariable)
	fmt.Printf("variable is of type: %T \n ", anatherVariable)

	// implicit types
	var website = "akashbiradar.com"
	fmt.Println(website)
	fmt.Printf("variable is of type: %T \n ", website)

	// no var style
	numberOfUsers := 1000
	fmt.Println(numberOfUsers)
	fmt.Printf("variable is of type: %T \n ", numberOfUsers)

	fmt.Println(LoginToken)
	fmt.Printf("variable is of type: %T \n ", LoginToken)

	var age int = 27             // Integer data type
	var salary float64 = 55000.5 // Floating-point data type
	name := "Akash"              // String data type (using type inference)

	fmt.Printf("Name: %s, Age: %d, Salary: %.2f\n", name, age, salary)
}
