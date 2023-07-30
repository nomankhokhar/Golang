package main

import "fmt"


	type User struct {
		Name string
		Email string
		Status bool
		Age int
	}

func main() {
	fmt.Println("Welcome to Structs")
	
	Noman := User{"Noman Ali", "younas@gmail.com",true,22}

	fmt.Println(Noman)
	fmt.Println(Noman.Name)
	fmt.Println(Noman.Email)
	fmt.Println(Noman.Age)
	fmt.Println(Noman.Status)
}