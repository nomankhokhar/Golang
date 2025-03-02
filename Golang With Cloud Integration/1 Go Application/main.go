package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// JSON to Struct Golang
type Words struct {
	Page  string   `json:"page"`
	Input string   `json:"input"`
	Words []string `json:"words"`
}

type MySlowReader struct {
	contents string
	pos      int
}

func (m *MySlowReader) Read(p []byte) (n int, err error) {
	if m.pos+1 <= len(m.contents) {
		n := copy(p, m.contents[m.pos:m.pos+1])
		m.pos++
		return n, nil
	}
	return 0, io.EOF
}

func main() {
	args := os.Args

	if len(args) < 2 {
		fmt.Println("Usage: ./http-get <argument>")
		os.Exit(1)
	}

	if _, err := url.ParseRequestURI(args[1]); err != nil {
		fmt.Printf("URL is in invalid format: %s\n", err)
		os.Exit(1)
	}
	response, err := http.Get(args[1])
	if err != nil {
		log.Fatal(err)
	}

	defer response.Body.Close()

	mySlowReaderInstance := &MySlowReader{
		contents: "Hello lol xD",
	}
	out, _ := io.ReadAll(mySlowReaderInstance)

	fmt.Println("Custome Struct and Interface", out)
	body, err := io.ReadAll(response.Body)

	if err != nil {
		log.Fatal(err)
	}
	if response.StatusCode != 200 {
		fmt.Printf("HTTP Status Code: %d\n Body: %s", response.StatusCode, body)
		os.Exit(1)
	}

	var words Words

	// convert request body data into Go Struct
	err = json.Unmarshal(body, &words)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("JSON Parsed\n Page: %s\nWords: %v\n", words.Page, strings.Join(words.Words, ", "))

}
