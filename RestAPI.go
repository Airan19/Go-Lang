// client.go
package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
)

func callServer() {
	// POST request
	postURL := "http://localhost:8080/post"
	postBody := []byte("Hello, this is a test message.")
	resp, err := http.Post(postURL, "text/plain", bytes.NewBuffer(postBody))
	if err != nil {
		fmt.Println("Error sending POST request:", err)
		return
	}
	defer resp.Body.Close()

	// Print response from POST request
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading POST response:", err)
		return
	}
	fmt.Println("POST Response:", string(respBody))

	// GET request
	getURL := "http://localhost:8080/get"
	resp, err = http.Get(getURL)
	if err != nil {
		fmt.Println("Error sending GET request:", err)
		return
	}
	defer resp.Body.Close()

	// Print response from GET request
	getRespBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading GET response:", err)
		return
	}
	fmt.Println("GET Response:", string(getRespBody))
}

func main() {
	// Start the server and listen for requests
	StartServer() // This will call the StartServer function from server.go
}
