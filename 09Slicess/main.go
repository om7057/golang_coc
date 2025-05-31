package main

import (
	"fmt"
	"slices"
	"sort"
)

func main() {
	fmt.Println("Welcome to slices in golang")

	var fruitList = []string{"Appe", "Banana", "Orange", "Grapes", "Pineapple"}
	fmt.Printf("Type of fruitList is %T \n", fruitList)

	fruitList = append(fruitList, "mango", "Watermelon")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:])
	fmt.Println(fruitList)

	fruitList = append(fruitList[:3])
	fmt.Println(fruitList)

	highScores := make([]int, 4)
	highScores[0] = 234
	highScores[1] = 131
	highScores[2] = 123
	highScores[3] = 346
	// highScores[4] = 346
	//Here it was giving error because we are trying to assign value to index 4 which is not present in the slice
	//But we can append values to the slice because it will automatically increase the size of the slice by reallocating the memory
	highScores = append(highScores, 456, 789)

	fmt.Println(sort.IntsAreSorted(highScores))
	sort.Ints(highScores)
	fmt.Println("highScores:", highScores)
	fmt.Println(sort.IntsAreSorted(highScores))

	//Removing a value from slice
	var courses = []string{"reactjs", "javascript", "python", "java", "swift", "c++"}
	fmt.Println(courses)

	var index int = 2 //Want the second index value to be removed
	// courses = append(courses[:index], courses[index+1:]...)
	courses = slices.Delete(courses, index, index+1)
	fmt.Println(courses)
}
