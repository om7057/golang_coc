package main

import "fmt"

func main() {
	fmt.Println("Welcome to maps in golang")


	// languages := make(map[int]string)
	languages := make(map[string]string)

	languages["JavaScript"]="JS"
	languages["Python"]="Py"
	languages["Java"]="Jv"
	languages["C++"]="Cpp"

	fmt.Println("List of all languages: \n", languages)
	fmt.Println("JavaScript short form is: ", languages["JavaScript"])
	
	delete(languages, "JavaScript")
	fmt.Println("List of all languages after deletion: \n", languages)


	//Looping through a map

	// for key, value := range languages {
	// 	fmt.Printf("For key %v, value is %v\n", key,value)
	// }

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n",value)
	}
}
