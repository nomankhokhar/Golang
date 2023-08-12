package main

import "fmt"

const (
	first = iota+1
	second
	third = 'a'
	forth 
)

func main() {
	fmt.Println("Welcome to function")
	Salaam()
	res := addTwoNum(3,3)
	fmt.Println(res)

	fmt.Println(preOrder(1,2,3,4,5,6))


	fmt.Println(testAdd(1,2))

	fmt.Println(first,second,third,forth)
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

func testAdd(num1 int , num2 int) (int , error){
	return num1 + num1 , nil
}