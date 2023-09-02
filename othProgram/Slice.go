package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	slice := make([]int, 0, 3) // Create an empty slice with initial capacity 3

	for {
		var input string
		fmt.Print("Enter an integer (or 'X' to quit): ")
		_, err := fmt.Scan(&input)

		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if input == "X" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter a valid integer or 'X' to quit.")
			continue
		}

		slice = append(slice, num)
		sort.Ints(slice) // Sort the slice in ascending order

		fmt.Println("Sorted slice:", slice)
	}
}
