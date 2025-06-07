package main

import (
	"fmt"
	"net/url"
)

// First handling the url
const myurl string = "https://pkg.go.dev"

func main() {
	fmt.Println("Web Request with URL package")
	fmt.Println(myurl)

	//Parsing the URL

	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)

	myurlpath := result.Path
	fmt.Println(myurlpath)

	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()
	fmt.Printf("Type of query params is %T\n ", qparams)

	fmt.Println(qparams["something"])

	for _, val := range qparams {
		fmt.Println("Param is :", val)
	}


	//Constructing a URL

	partsofurl := &url.URL{
		Scheme: "https",
		Host:   "lco.dev",
		Path: "learn",
		RawPath: "user=hitesh",
	}

	anotherurl := partsofurl.String()
	fmt.Println("Constructed URL is:", anotherurl)
}
