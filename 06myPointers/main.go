package main

import (
	"fmt"
	"strings"
)

func swap(m1, m2 *int) {
	var temp int
	temp = *m1
	*m1 = *m2
	*m2 = temp
}

func main() {
	m1, m2 := 2, 3
	fmt.Println(m1, m2)
	swap(&m1, &m2)
	fmt.Println(m1, m2)

	mone := "my name"
	Chand := "name"
	fmt.Println(strings.Split(mone, " "), mone+Chand)

	fmt.Println("MyPointers")
	numBer := 200
	var ptr *int = &numBer
	numBer = 210
	*ptr = 200
	fmt.Println(*ptr)
}
