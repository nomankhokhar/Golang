package main

import "fmt"

const LoginToken = "Noman Ali"

func main() {

	var username string = "Noman Ali"
	fmt.Println("String", username)
	fmt.Printf("%T\n", username)

	var isloggedIn bool = false
	fmt.Println("Bool", isloggedIn)
	fmt.Printf("%T\n", isloggedIn)

	var smallVal uint8 = 255
	fmt.Println("Unsigned 8 Bit Value", smallVal)
	fmt.Printf("%T\n", smallVal)

	var floating float64 = 255.44554455
	fmt.Println("Floating Point Value", floating)
	fmt.Printf("%T\n", floating)

	var anotherVeriable int
	fmt.Println("Integer Value ", anotherVeriable)
	fmt.Printf("%T\n", anotherVeriable)

	numOfUser := 300
	fmt.Println("Num of User", numOfUser)

	fmt.Println(LoginToken)
}
