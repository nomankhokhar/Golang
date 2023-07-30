package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to my Time")


	presentTime := time.Now()

	fmt.Println(presentTime)
	createdTime := presentTime.Format("01-02-2006")
	fmt.Println(createdTime)
	fmt.Println(time.Date(2023, time.August, 10 , 23, 23, 0 , 0,time.UTC))

}