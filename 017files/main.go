package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Velkome to File Reading")

	content := "Noman Ali is a Nil Student"

	file, err := os.Create("./nomanFile.txt")
	checkNilError(err)
	defer file.Close()

	length, err := io.WriteString(file, content)
	checkNilError(err)

	fmt.Println("Data length is : ", length)

	readFile("./nomanFile.txt")
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) {
	dataBytes, err := ioutil.ReadFile(filename)
	checkNilError(err)

	fmt.Println("File data is : ", string(dataBytes))
}
