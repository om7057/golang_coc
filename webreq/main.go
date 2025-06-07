package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Welcome to the Go HTTP Server!")
	// PerformGetRequest()
	// PerformPostJsonRequest()
	PerformPostFromRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:8000/get"

	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status Code:", response.StatusCode)
	fmt.Println("Content Length:", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)
	byteCount, _ := responseString.Write(content)

	fmt.Println("Byte count is :", byteCount)

	fmt.Println(responseString.String())

	// fmt.Println(string(content))
}

func PerformPostJsonRequest() {
	const myurl = "http://localhost:8000/post"

	//fake json payload

	requestbody := strings.NewReader(`
	{
		"coursename":"golang",
		"price":0,
		"platform":"google.in"
	}	
	`)
	response, err := http.Post(myurl, "application/json", requestbody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFromRequest() {
	const myurl = "http://localhost:8000/postform"

	//form-data

	data := url.Values{}
	data.Add("firstname", "om")
	data.Add("lastname", "kulkarni")
	data.Add("email", "om.kulkarni@walchandsangli.ac.in")

	response, err := http.PostForm(myurl, data)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}
