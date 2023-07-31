package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")

	// days := []string{"Sun", "Mon","Tue","Wed","Thur","Fri","Sat"}

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days[i])
	// }

	// for i := range days {
	// 	fmt.Println(days[i])
	// }

	// for index ,value := range days {
	// 	fmt.Println(index, value)
	// }

	rougueValue := 1

	for rougueValue < 10 {

		if rougueValue == 5 {
			rougueValue++
			goto loc
		}
		break
		fmt.Println(rougueValue)
		rougueValue++
	}
loc:
	fmt.Println("Loc lable is here")

}