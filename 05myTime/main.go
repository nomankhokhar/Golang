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

	fmt.Println(time.Date(
		2026,        // Year
		time.August, // Month
		10,          // Day
		23,          // Hour
		23,          // Minute
		0,           // Second
		0,           // Nanosecond
		time.UTC,    // Timezone
	))
}
