package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in golang")


	days := []string{"Sunday", "Tuestday", "Wednesday", "Friday", "Saturday"}

	fmt.Println(days)	

	// for d :=0 ; d<len(days);d++{
	// 	fmt.Println("day is: ", days[d])
	// }


	// for i := range days{
	// 	fmt.Println(days[i])
	// }


	// for idx, day := range days {
	// 	fmt.Printf("Index is %v and index is %v\n", idx, day)
	// }

	for day := range days {
		fmt.Printf("Day is %v \n", day)
	}

	rogueValue := 1

	for rogueValue < 10 {
		// if rogueValue == 5 {
		// 	break
		// }
		if rogueValue == 2 {
			goto google
		}
		if rogueValue == 5 {
			rogueValue++
			continue
		}

		fmt.Println("Value is :\n", rogueValue)
		rogueValue++
	}


	google : 
	     fmt.Println("Jumping at google.com")
}
