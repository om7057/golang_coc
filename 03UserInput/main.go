package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "Welcome to user input"
	fmt.Println(welcome)


	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our pizza: ")
	
	//comma ok syntax || comma error syntax

	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating us: ", input)
	fmt.Printf("Type of the rating that you gave is : %T", input)
}