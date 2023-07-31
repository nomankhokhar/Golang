package main

import "fmt"


	type User struct {
		Name string
		Email string
		Status bool
		Age int
	}

	func (u User)GetStatus(){
		fmt.Println("Is user active : ", u.Status)	
	}

	func (u User)GetNewMail(){
		u.Email = "a@gmail.com"
		fmt.Println("New mail is :", u.Email)
	}

func main() {
	fmt.Println("Welcome to Structs")
	
	Noman := User{"Noman Ali", "younas@gmail.com",true,22}

	fmt.Println(Noman)
	fmt.Println(Noman.Name)
	fmt.Println(Noman.Email)
	fmt.Println(Noman.Age)
	fmt.Println(Noman.Status)

	Noman.GetStatus()
	Noman.GetNewMail()
}



