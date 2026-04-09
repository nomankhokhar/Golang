package main

import "fmt"

// Defer works like a Stack (LIFO - Last In First Out) that executes deferred functions in reverse order.

func main() {
	fmt.Println("Welcome to Defer Golang")

	defer fmt.Println("One")
	defer fmt.Println("World")
	defer fmt.Println("Hello")

	defferLast()
}

func defferLast() {
	for i := 0; i < 10; i++ {
		defer fmt.Println(i + 1)
	}
}
