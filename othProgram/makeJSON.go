package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string

	fmt.Print("Enter a name: ")
	fmt.Scanln(&name)

	fmt.Print("Enter an address: ")
	fmt.Scanln(&address)

	data := make(map[string]string)
	data["name"] = name
	data["address"] = address

	jsonData, err := json.Marshal(data)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("JSON Object:")
	fmt.Println(string(jsonData))
}
