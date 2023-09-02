package main

import (
	"fmt"
	"strings"
)

func main() {
	var word string
	fmt.Scanln(&word)
	word = strings.ToLower(word)
	word = strings.TrimSpace(word)
	fmt.Printf(word)
	if strings.HasPrefix(word, "i") && strings.Contains(word, "a") && strings.HasSuffix(word, "n") {
		fmt.Printf("Found")
	} else {
		fmt.Printf("Not Found")
	}
}
