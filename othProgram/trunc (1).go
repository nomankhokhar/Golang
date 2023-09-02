package main

import (
	"fmt"
	"strconv"
)

func main() {
	var in float64
	fmt.Printf("please enter a float: ")
	fmt.Scan(&in)
	x := int(in)
	fmt.Printf("your input truncated: ")
	fmt.Printf(strconv.Itoa(x) + "\n")
	return
}
