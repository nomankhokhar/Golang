package main

import "fmt"

func main() {
	fmt.Println("Enter Floating Point No. ")
	var first float64
	fmt.Scanln(&first)
	fmt.Printf("The truncated number is %.f", first)
}
