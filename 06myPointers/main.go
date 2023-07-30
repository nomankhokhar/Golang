package main

import "fmt"

func main() {

	fmt.Println("MyPointers")
	numBer := 200
	var ptr *int = &numBer
	numBer = 210
	*ptr = 200
	fmt.Println(*ptr)

}