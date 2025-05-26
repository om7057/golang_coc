package main

import (
	"fmt"
	"math/big"
	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to maths in golan")

	// var numberOne int = 2;
	// var numberTwo float64 = 4
	// fmt.Println("Sum is: ", numberOne + numberTwo)

	//Random Number
	// r := rand.New(rand.NewSource(time.Now().UnixNano()))
	// fmt.Println("random no. ", r.Intn(5))

	//Random from crypto

	myRandomNumber, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNumber)
}
