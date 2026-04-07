package main

import "fmt"

func main() {
	fmt.Println("Welcome to IfElse")

	loginCount := 2

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "Something Else"
	}

	fmt.Println(result)

	if num := 3; num < 10 && num > 1 {
		fmt.Println("Num is between 1 and 10")
	}

	// This is Slice
	arr := []string{"Noman", "Ali", "JS", "Rust", "Golang"}

	for index, value := range arr {
		fmt.Println(index, value)
	}

	// This is Map
	myMap := make(map[string]interface{})
	myMap["name"] = "Noman"
	myMap["age"] = 30
	myMap["isAdmin"] = true

	for index, value := range myMap {
		fmt.Println(index, value)
	}
}
