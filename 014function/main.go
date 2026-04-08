package main

import "fmt"

// iota will increment by 1 for each constant in the block
const (
	first  = iota + 1 // This will add 1 to the iota value
	second = 65
	third  = 'a'
	forth  = 0
)

func main() {
	fmt.Println("welcome to function")

	Salaam()
	fmt.Println(preOrder(1, 2, 3, 4, 5, 6, 7))
	fmt.Println(testAdd(1, 2))
	fmt.Println(first, second, third, forth)
}

func Salaam() {
	fmt.Println("Velkome Sir")

	res := addTwoNum(3, 3)
	fmt.Println(res)
}

func addTwoNum(a, b int) int {
	return a + b
}

func preOrder(values ...int) int {
	Sumvalue := 0
	for _, value := range values {
		fmt.Println(value)
		Sumvalue += value
	}
	fmt.Println("Sum:", Sumvalue)
	return Sumvalue
}

func testAdd(a, b int) int {
	fmt.Println(a, b)
	return a + b
}
