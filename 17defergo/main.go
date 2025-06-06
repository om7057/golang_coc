package main

import "fmt"

func main() {
	defer fmt.Println("Golang in defer ")
	fmt.Println("Defer in")
	fmt.Println("Defer in")
	defer fmt.Println("Defer in g")
	fmt.Println("Defer in go")
	defer fmt.Println("Defer in gol")
	fmt.Println("Defer in gola")
	fmt.Println("Defer in golan")
}
