// Make server using gojin
// take first name and lastname then make struck return json responce POST Requence
// POST API AND GET API
package main

import "fmt"

func main()  {
	var a interface{}
	a = true


	var b *int
	var c *int

	b = ptr(10)
	c = ptr(10)
	
	fmt.Println(b, c)

	switch t:= a.(type) {
	case int64:
		fmt.Println("Type is int : ", t)
	case float64:
		fmt.Println("Type is float : ", t)
	case string:
		fmt.Println("Type is string : ", t)
	case bool:
		fmt.Println("Type is bool : ", t)
	default:
		fmt.Println("Default")
	}
}

func ptr(ptr int)(*int){
	return &ptr
}