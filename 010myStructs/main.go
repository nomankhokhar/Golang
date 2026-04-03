package main

import (
	"fmt"
)

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type Car interface {
	Drive()
	Stop()
}

type Lambo struct {
	LamboModel string
}

func NewModelLambo(arg string) Car {
	return &Lambo{arg}
}

func (c *Lambo) Stop() {
	fmt.Println("Stopping Lambo:", c)
}

func (c *Lambo) Drive() {
	fmt.Println("Driving Lambo:", c)
}

type Chevy struct {
	ChevyModel string
	Name       string
}

func NewModelChevy(arg string, name string) Car {
	return &Chevy{arg, name}
}

func (c *Chevy) Stop() {
	fmt.Println("Stopping Chevy:", c)
}

func (c *Chevy) Drive() {
	fmt.Println("Driving Chevy:", c)
}

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("Welcome to Structs")

	Noman := User{"Noman Ali", "younas@gmail.com", true, 22}
	fmt.Println(Noman)

	Nomi := User{}

	Nomi.Name = "Nomi Bhai"
	Nomi.Email = "nomi@gmail.com"
	Nomi.Status = true
	Nomi.Age = 25

	fmt.Println(Nomi)

	l := NewModelLambo("KaliNoman")
	c := NewModelChevy("C986", "NomanAli")

	n := Lambo{"Lambda"}
	n.LamboModel = "Ali"
	fmt.Println(n.LamboModel)
	n.Drive()

	l.Drive()
	c.Drive()

	fmt.Println("vim-go")
	Anything(2.44)
	Anything("Angad")
	Anything(struct{}{})

	MapExample()
}
