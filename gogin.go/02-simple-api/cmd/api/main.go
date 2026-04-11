package main

import (
	"fmt"
	"simple-api/model"
	"simple-api/routes"
)

var data model.Api

func main() {
	fmt.Println("Welcome to Golang->API")
	r := routes.ApiRoute()
	r.Run(":3000")
}
