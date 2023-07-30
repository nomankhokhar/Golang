package main

import (
	"bufio"
	"fmt"
	"os"
)

func main(){
	welcome := "Noman Ali Velkome you to user input"

	fmt.Println("Enter the rating input for Pizza : ")
	reader := bufio.NewReader((os.Stdin))
	
	fmt.Println(welcome)

	// comma ok || err

	input , _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating" , input)
	fmt.Printf("Thanks for rating and type of rating %T" , input)
}