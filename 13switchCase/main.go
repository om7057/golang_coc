package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to switch case in golang")

	// rand.Seed(time.Now().UnixNano())
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	diceNumber := r.Intn(6) + 1
	fmt.Println("Value on the dice is :\n", diceNumber)


	switch diceNumber {
	case 1:
		fmt.Println("Dice value is 1 , you can open")
		fallthrough
	case 2:
		fmt.Println("Dice value is 2 , you can move 2 steps")
		fallthrough
	case 3:
		fmt.Println("Dice value is 3 , you can move 3 steps")
		fallthrough
	case 4:
		fmt.Println("Dice value is 4 , you can move 4 steps")
		fallthrough
	case 5:
		fmt.Println("Dice value is 5 , you can move 5 steps")
		fallthrough
	case 6:
		fmt.Println("Dice value is 6 , you can move 6 steps and roll again")
		
	default:
		fmt.Println("What the hell is this ?")
	}
}
