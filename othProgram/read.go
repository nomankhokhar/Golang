package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Name struct {
	Fname string
	Lname string
}

func main() {
	var filename string

	fmt.Print("Enter the name of the text file: ")
	fmt.Scanln(&filename)

	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer file.Close()

	names := make([]Name, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		name := Name{
			Fname: fields[0],
			Lname: fields[1],
		}

		names = append(names, name)
	}

	if scanner.Err() != nil {
		fmt.Println("Error:", scanner.Err())
		return
	}

	fmt.Println("Names in the file:")
	for _, name := range names {
		fmt.Printf("%s %s\n", name.Fname, name.Lname)
	}
}
