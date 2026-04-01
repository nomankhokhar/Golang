package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Golang")

	// This will create an array of 4 strings We have to manually assign values
	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Kela"
	fruits[3] = "Saib"

	fmt.Println("Total fruits are  : ", fruits)
	fmt.Println("Total fruits are and length : ", len(fruits))

	// This will create an array of 3 strings
	var vegList = [3]string{}
	vegList[0] = "Potato" // Other will set to Empty

	// This will create an array of 4 strings
	var VegList = [...]string{"Carrot", "Broccoli", "Spinach", "Cucumber"}

	fmt.Println(vegList)
	fmt.Println(VegList)
}
