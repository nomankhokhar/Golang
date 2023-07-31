package main

import "fmt"

func main() {
	fmt.Println("Welcome to function")
	Salaam()
	res := addTwoNum(3,3)
	fmt.Println(res)

	fmt.Println(preOrder(1,2,3,4,5,6))
}


func preOrder (value ...int) (string,int) {
	total := 0
	for _ , val := range value {
		total += val
	} 
	return "The total is :",total
}

func addTwoNum (a int, b int) int {	
	return (a + b)
}

func Salaam(){
	fmt.Println("VelKome Sir")
}