package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: ./hell-world <argument>\n")
		os.Exit(1)
	}

	fmt.Println("Hello Cloud", args[1:])
}
