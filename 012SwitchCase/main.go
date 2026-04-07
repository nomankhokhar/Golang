package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to SwitchCase")

	// This seed will ensure that we get a different random number each time
	rand.Seed(time.Now().UnixNano())
	// This Intn(10) will generate a random number between 0 and 9
	diceNumber := rand.Intn(10)

	fmt.Println("Value of dice is", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("Value of dice is ", diceNumber)
	case 2:
		fmt.Println("Value of dice is ", diceNumber)
	case 3:
		fmt.Println("Value of dice is ", diceNumber)
	case 4:
		fmt.Println("Value of dice is ", diceNumber)
	case 5:
		fmt.Println("Value of dice is ", diceNumber)
	case 6:
		fmt.Println("Value of dice is ", diceNumber)
	default:
		fmt.Println("Default Value of dice is ", diceNumber)
	}
}
