package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("If else in goland")

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the value for login count: ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // Remove any trailing newline or spaces

	loginCount, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Invalid input, please enter a number")
		return
	}

	fmt.Println("Your entered value is: ", input)

	var result string

	if loginCount < 10 {
		result = "Regular user"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Exactly 10 logins"
	}

	fmt.Println("Result is: ", result)

	if 9%2 == 0 {
		fmt.Println("9 is even")
	} else {
		fmt.Println("9 is odd")
	}

	//Mostly for handling web requests and serving them on the go 

	if num := 3; num < 10 {
		fmt.Println("Number is less than 10")
	} else {
		fmt.Println("Number is not less than 10")
	}
}
