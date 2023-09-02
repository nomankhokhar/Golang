package main

import "fmt"

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

func NewModal(arg string) Car {
	return &Lambo{arg}
}

type Chevy struct {
	ChevyModel string
}

// You can Add Pointer Address and Pointing But ByDefault Go Pick it up Automatically

func (l *Lambo) Stop() {
	fmt.Println("Stopping lambo")
}

func (l *Lambo) Drive() {
	fmt.Println("Lambo on the move")
	fmt.Println(l.LamboModel)
}

func (c *Chevy) Drive() {
	fmt.Println("Chevy on the move")
	fmt.Println(c.ChevyModel)
}

func Anything(anything interface{}) {
	fmt.Println(anything)
}

func main() {
	fmt.Println("Welcome to Structs")

	Noman := User{"Noman Ali", "younas@gmail.com", true, 22}

	Nomi := User{}

	Nomi.Name = "Noman"

	fmt.Println(Nomi)
	fmt.Println(Noman)
	fmt.Println(Noman.Name)
	fmt.Println(Noman.Email)
	fmt.Println(Noman.Age)
	fmt.Println(Noman.Status)

	l := NewModal("KaliNoman")
	c := Chevy{"C986"}

	l.Drive()
	c.Drive()

	// Interface is Usefull When YOu not know about type of the Data

	fmt.Println("vim-go")
	Anything(2.44)
	Anything("Angad")
	Anything(struct{}{})

}

// Delete Key in the Map using Delete Value

// package main

// import "fmt"

// func main() {
// 	// Create an instance of an empty struct
// 	emptyStruct := struct{}{}

// 	// Use the empty struct, for example, as a placeholder in a map
// 	myMap := make(map[string]struct{})
// 	myMap["key1"] = emptyStruct
// 	myMap["key2"] = emptyStruct

// 	// Check if a key exists in the map
// 	if _, exists := myMap["key1"]; exists {
// 		fmt.Println("Key1 exists in the map.")
// 	}

// 	// Delete a key from the map
// 	delete(myMap, "key2")

// 	// Check if a key exists after deletion
// 	if _, exists := myMap["key2"]; !exists {
// 		fmt.Println("Key2 does not exist in the map.")
// 	}
// }
