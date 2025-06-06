package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in golang")

	content := "This needs to go in a file - LCO.in"

	file, err := os.Create("./mylcogofile.txt")

	if err != nil {
		panic(err) //Shuts the program and prints the error
	}

	length, err := io.WriteString(file, content)
	// if err != nil {
		// panic(err) //Shuts the program and prints the error
	// }
	checkNilErr(err) 
	
	fmt.Println("Length is ", length)

	defer file.Close()

	readFile("./mylcogofile.txt")
}



//For reading a file and doing something with it we have a separate utility given to us
func readFile (filename string) {
	
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)
	fmt.Println("Text data inside file is \n" , string(databyte) )

}


func checkNilErr(err error){
	if err != nil {
		panic(err)
	}
}