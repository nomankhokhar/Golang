package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	// This is Standard For Loop
	for index := 0; index < len(days); index++ {
		fmt.Println(days[index])
	}

	// This is Go Style For Loop
	for index, value := range days {
		fmt.Println(index, value)
	}

	rougueValue := 1

	for rougueValue < 10 {
		if rougueValue == 5 {
			rougueValue++
			goto loc
		}
		fmt.Println(rougueValue)
		rougueValue++

	}

loc:
	fmt.Println("Loc label is here", rougueValue)
}
