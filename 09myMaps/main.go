package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps of User")

	language := make(map[string]string)

	language["jais"] = "Ali"
	language["ruby"] = "Noman Ali"
	language["java"] = "Noman Ali Khan"
	language["khan"] = "Khan Lala Famous"

	fmt.Println(language)
	fmt.Println(language["js"])

	delete(language , "ruby")

	fmt.Println(language)


	for key, value := range language{
		fmt.Println(key, " : " , value)
	}
}