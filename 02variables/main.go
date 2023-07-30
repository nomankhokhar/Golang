package main

import "fmt"

const LoginToken = "Noman Ali"

func main() {

	var username string = "Noman Ali"
	fmt.Println(username)
	fmt.Printf("%T\n", username) 



	var isloggedIn bool = false
	fmt.Println(isloggedIn)
	fmt.Printf("%T\n", isloggedIn)
	
	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("%T\n", smallVal)


	var floating float64 = 255.44554455
	fmt.Println(floating)
	fmt.Printf("%T\n", floating)


	var anotherVeriable int 
	fmt.Println(anotherVeriable)
	fmt.Printf("%T\n", anotherVeriable)

	numOfUser:= 300
	fmt.Println(numOfUser)

	fmt.Println(LoginToken)
}