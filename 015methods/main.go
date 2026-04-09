package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u *User) GetStatus() {
	fmt.Println(u.Status)
}

func (u *User) GetNewMail() {
	u.Name = "Ali"
	fmt.Println(u.Name)
}

func main() {
	fmt.Println("Welcome to Structs")

	Noman := User{"Noman Ali", "younas@gmail.com", true, 22}

	fmt.Println(Noman)
	fmt.Println(Noman.Name)
	fmt.Println(Noman.Email)
	fmt.Println(Noman.Age)
	fmt.Println(Noman.Status)

	Noman.GetStatus()
	Noman.GetNewMail()

	fmt.Println(Noman)
}
