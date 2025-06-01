package main

import "fmt"

func main() {
	fmt.Println("Welcome to functions in golang")
	result := adder(3, 5)
	fmt.Println("Result is:", result)
	greeter()

	proResult, myMessage := proAdder(1, 2, 3, 4, 5)
	fmt.Println("Pro Result is:", proResult)
	fmt.Println("Pro Message is:", myMessage)
}

func adder (valOne int, valTwo int) int {
	return valOne + valTwo
}

func proAdder(values ...int) (int, string) {
	total := 0
	for _, val := range values {
		total += val
	}

	return total, "Hi from proAdder"

}

func greeter() {
	fmt.Println("Namastey from golang")
}
