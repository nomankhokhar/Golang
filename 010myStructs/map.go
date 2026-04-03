package main

import (
	"fmt"
)

func MapExample() {
	fmt.Println("Welcome to Map go")
	// Create an instance of an empty struct
	emptyStruct := struct {
		noman string
	}{
		noman: "noman",
	}

	// Use the empty struct, for example, as a placeholder in a map
	myMap := make(map[string]struct {
		noman string
	})

	myMap["key1"] = emptyStruct
	myMap["key2"] = emptyStruct

	fmt.Println("Map contents:", myMap)

	if _, exists := myMap["key1"]; exists {
		fmt.Println("Key1 exists in the map")
	}

	value, exists := myMap["key2"]

	fmt.Println("key2 value:", value)
	fmt.Println("key2 exists:", exists)

	// Delete a key from the map
	delete(myMap, "key2")

	// check if key exists after deletion
	if _, exists := myMap["key2"]; !exists {
		fmt.Println("key2 does not exist in the map.")
	}
}
