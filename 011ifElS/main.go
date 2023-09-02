package main

import "fmt"

func main() {
	fmt.Println("Welcome to If Else")

	loginCount := 2

	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else {
		result = "Something Else"
	}

	fmt.Println(result)
	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	}

	arr := []string{"Nomam", "Ali", "JS", "Rust", "Golang"}

	for index, val := range arr {
		fmt.Println(index, val)
	}

	myMap := make(map[string]interface{})
	myMap["name"] = "Noman"
	myMap["age"] = 21
	myMap["job"] = "CS"

	for index, value := range myMap {
		fmt.Println(index, value)
	}

}
