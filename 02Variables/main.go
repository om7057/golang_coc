package main

import "fmt"


// jwtToken := 300000 -> Not allowed outside function

const LoginToken string = "qwertyuiop"

func main() {
	// fmt.Println("Variables")
	var username string = "Om" 
	fmt.Println(username)
	fmt.Printf("Variable is of type : %T \n",username)


	var isLoggedIn bool = true 
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable is of type : %T \n",isLoggedIn)



	var smallVal uint8  = 255
	fmt.Println(smallVal)
	fmt.Printf("Variable is of type : %T \n",smallVal)


	var smallFloat float64 = 255.455554542
	fmt.Println(smallFloat)
	fmt.Printf("Variable is of type : %T \n",smallFloat)

	//Default values and aliases
	var anotherVariable int 
	fmt.Println(anotherVariable)
	fmt.Printf("Variable is of typr : %T \n", anotherVariable)

	//Implicit type
	var website = "google.com"
	fmt.Println(website)

	//No var style 
	numberOfUser := 300000
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n",LoginToken)
}