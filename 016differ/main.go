package main

import "fmt"

func main() {
	fmt.Println("Welcome to Differ")



	defer fmt.Println("One")
	defer fmt.Println("World")
	defer fmt.Println("Hello")

	defferLast()
}

func defferLast(){
	for i := 0 ; i < 10; i++ {
		defer fmt.Println(i+1)
	}
}