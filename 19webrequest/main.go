package main

import (
	"fmt"
	"io"
	"net/http"
)
const url = "https://pkg.go.dev"

func main() {
	fmt.Println("LCO Web Request")
	response, err :=http.Get(url)
	if err!=nil{
		panic(err)
	}
	fmt.Printf("Type of response is %T\n",response)

	defer response.Body.Close() //Caller's responsibility to close the connection


	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Println(content)
}
