package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	name        string
	description string
	phone       int
}

func main() {

	fmt.Println("JSON VALUES")

	// var wrapData map[string]string      // this will take only string value in both key and value
	var wrapData map[string]interface{} // this will take only string value in key but tak any value in value
	// wrapData := make(map[string]interface{})

	data := `{"name": "Noman Ali", "description" : 54554 , "phone" : 3122652}`

	err := json.Unmarshal([]byte(data), &wrapData)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(wrapData)

	val, ok := wrapData["name"]

	if ok {
		fmt.Println(val)
	}
	// Using Struct SAve JSON DATA

	var wrapDataJSON User

	errStr := json.Unmarshal([]byte(data), &wrapDataJSON)

	if errStr != nil {
		fmt.Println(errStr)
	}

	fmt.Println(wrapData)

	// Marshall vs UnMarshall

	byteData, err := json.Marshal(&wrapData)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(byteData)
	fmt.Println(string(byteData))

	err = json.Unmarshal(byteData, &wrapData)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(wrapData)

	// Marshell the Data make json

	type Person struct {
		Name string
		Age  int
	}

	person := Person{Name: "NomiBhai", Age: 22}

	jsonData, err := json.Marshal(person)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(jsonData))

	// UnMarshalling the DATA to store data into MAP or Struct

	jsonDataMar := []byte(`{"description":54554,"name":"Noman Ali","phone":3122652}`)
	var personMar Person

	errMar := json.Unmarshal(jsonDataMar, &personMar)

	if err != nil {
		fmt.Println(errMar)
	}

	fmt.Println(personMar)
}
