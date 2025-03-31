package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LoginRequest struct {
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

func DoLoginRequest(url string, password string) (string, error) {
	loginRequest := LoginRequest{
		Password: password,
	}

	body, err := json.Marshal(loginRequest)
	if err != nil {
		return "", fmt.Errorf("marshal error %s", err)
	}

	response, err := http.Post(url, "application/json", bytes.NewBuffer(body))
	if err != nil {
		return "", fmt.Errorf("http Post error: %s", err)
	}

	defer response.Body.Close()

	mySlowReaderInstance := &MySlowReader{
		contents: "Hello lol xD",
	}
	out, _ := io.ReadAll(mySlowReaderInstance)

	fmt.Println("Custome Struct and Interface", out)
	resBody, err := io.ReadAll(response.Body)

	if err != nil {
		return "", fmt.Errorf("ReadAll error: %s", err)
	}

	if response.StatusCode != 200 {
		return "", fmt.Errorf("HTTP Status Code: %d\n Body: %s", response.StatusCode, string(resBody))
	}

	var page Page

	err = json.Unmarshal(resBody, &page)
	if err != nil {
		return "", &RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(body),
			Err:      fmt.Sprintf("Page unmarshal error: %s", err),
		}
	}

	var loginResponse LoginResponse

	// convert request body data into Go Struct
	err = json.Unmarshal(resBody, &loginResponse)

	if err != nil {
		return "", &RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(resBody),
			Err:      fmt.Sprintf("Unmarshal error: %s", err),
		}
	}

	if loginResponse.Token == "" {
		return "", &RequestError{
			HTTPCode: response.StatusCode,
			Body:     string(resBody),
			Err:      "Empty token replied",
		}
	}

	return loginResponse.Token, nil
}
