package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
)

// Func Interface
type Response interface {
	GetResponse() string
}

type Page struct {
	Name string `json:"page"`
}

type Words struct {
	Input string   `json:"input"`
	Words []string `json:"words"`
}
type Occurance struct {
	Words map[string]int `json:"words"`
}

func (w Words) GetResponse() string {
	return strings.Join(w.Words, ", ")
}

func (o Occurance) GetResponse() string {
	out := []string{}
	for word, occurance := range o.Words {
		out = append(out, fmt.Sprintf("%s (%d)", word, occurance))
	}
	return strings.Join(out, ", ")
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
	res, err := doRequest(args[1])

	if err != nil {
		if requestError, ok := err.(*RequestError); ok {
			fmt.Printf("Error: %s (HTTP Code %d, Body: %s)\n", requestError.Err, requestError.HTTPCode, requestError.Body)
		}
		os.Exit(1)
	}

	if err == nil {
		fmt.Println("No response")
		os.Exit(1)
	}

	fmt.Printf("Response %s", res.GetResponse())

}

func doRequest(requestURL string) (Response, error) {
	if _, err := url.ParseRequestURI(requestURL); err != nil {
		return nil, fmt.Errorf("validation error: %s", err)
	}
	response, err := http.Get(requestURL)
	if err != nil {
		return nil, fmt.Errorf("http Get error: %s", err)
	}

	defer response.Body.Close()

	mySlowReaderInstance := &MySlowReader{
		contents: "Hello lol xD",
	}
	out, _ := io.ReadAll(mySlowReaderInstance)

	fmt.Println("Custome Struct and Interface", out)
	body, err := io.ReadAll(response.Body)

	if err != nil {
		return nil, fmt.Errorf("ReadAll error: %s", err)
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("HTTP Status Code: %d\n Body: %s", response.StatusCode, body)
	}

	var page Page

	err = json.Unmarshal(body, &page)
	if err != nil {
		return nil, &RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Sprintf("Page unmarshal error: %s", err),
		}
	}

	// convert request body data into Go Struct
	err = json.Unmarshal(body, &page)

	if err != nil {
		return nil, &RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Sprintf("Unmarshal error: %s", err),
		}
	}

	switch page.Name {
	// localhost:300/words
	// {
	// 	"page": "words",
	// 	"input": "hello world hello",
	// 	"words": ["hello", "world", "hello"]
	// }
	case "words":
		var words Words
		err = json.Unmarshal(body, &words)
		if err != nil {
			return nil, &RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Words Unmarshal error: %s", err),
			}
		}
		return words, nil

		// localhost:300/occurrence
	// {
	// 	"page": "occurrence",
	// 	"words": {
	// 		"hello": 2,
	// 		"world": 1
	// 	}
	// }
	case "occurrence":
		var occurrence Occurance
		err = json.Unmarshal(body, &occurrence)
		if err != nil {
			return nil, &RequestError{
				HTTPCode: response.StatusCode,
				Body:     string(body),
				Err:      fmt.Sprintf("Occurrence Unmarshal error: %s", err),
			}
		}
		return occurrence, nil
	}
	return nil, nil
}
