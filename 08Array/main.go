package main

import "fmt"

func main() {
	fmt.Println("Welcome to the arrays in golang")


	var fruitList = [4]string{}

	fruitList[0] = "Apple"
	// fruitList[1] = "Orange"
	fruitList[2] = "Banana"
	fruitList[3] = "Mango"
	
	fmt.Println("Fruit list is :", fruitList)
	fmt.Println("Fruit list is :", len(fruitList))

	var vegList = [3]string{"Potato", "Tomato", "Onion"}
	fmt.Println("Vegy list is :", vegList)
	fmt.Println("Vegy list is :", len(vegList))
}
