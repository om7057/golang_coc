package main

import "fmt"

func main() {
	fmt.Println("Welcome to pointers in Go!")


	// var ptr *int 
	// fmt.Println(ptr)

	myNumber := 23 

	var ptr *int = &myNumber // & gives the address of myNumber
	fmt.Println("Value of the pointer is ",ptr)
	fmt.Println("Value of the pointer is ",*ptr)
	

	*ptr = *ptr * 2 
	fmt.Println("myNumber is now ", myNumber)
}
