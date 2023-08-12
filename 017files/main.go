package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Velcome to Files\n")

	content := "Noman Ali is a Nil Student"

	file , err := os.Create("./nomanFile.txt")
	checkNilError(err)

	defer file.Close()

	lenght , err := io.WriteString(file, content)
	checkNilError(err)
	fmt.Println("Data lenght is : ", lenght)
	readFile("./nomanFile.txt")
}

func checkNilError (err error){
	if err != nil {
		panic(err)
	}
}

func readFile(filename string){
	dataBytes , err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}

	fmt.Println("File data is : " , string(dataBytes))
}