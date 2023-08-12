package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array in Golang")

	var fruits [4]string

	fruits[0] = "Apple"
	fruits[1] = "Banana"
	fruits[2] = "Kela"
	fruits[3] = "Saib"

	fmt.Println("Total fruits are : ", fruits)
	fmt.Println("Total fruits are and length : ", len(fruits))

	var vegList = [3]string{"Potato", "Beans", "mushroom"}

	fmt.Println(vegList)
}
