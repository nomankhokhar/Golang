package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to SwitchCase")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6)+1
	fmt.Println("Value of dice is ", diceNumber)



	switch diceNumber {
	case 1: fmt.Println("Value of dice is ", diceNumber)
	case 2:	fmt.Println("Value of dice is ", diceNumber)
	case 3:	fmt.Println("Value of dice is ", diceNumber)
	case 4:	fmt.Println("Value of dice is ", diceNumber)
	case 5:	fmt.Println("Value of dice is ", diceNumber)
	case 6:	fmt.Println("Value of dice is ", diceNumber)
	}
}