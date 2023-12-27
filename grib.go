package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"os"

	"github.com/ncruces/zenity"
)

// selectGrib selects a file using a file dialog and calls the gribRequest function with the selected file.
//
// No parameters.
// No return type.
func selectGrib() {
	f, err := zenity.SelectFile(
		zenity.Filename("./"),
		zenity.FileFilters{
			{
				Name:     "Images",
				Patterns: []string{"*.png", "*.jfif", "*.jpeg", "*.jpg"},
				CaseFold: true,
			},
		},
	)

	if err != nil {
		fmt.Println("\n please select a valid image")
		panic(err)
	}

	gribRequest(f)

}

// gribRequest sends an HTTP request with a file attached to the given URL using the multipart form data format.
//
// Parameters:
// - f: the file path to be sent in the request.
//
// Returns: None.
func gribRequest(f string) {
	// Open the file
	file, err := os.Open(f)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a buffer to store the file contents
	requestBody := &bytes.Buffer{}
	writer := multipart.NewWriter(requestBody)

	// Create a form field for the file
	fileField, err := writer.CreateFormFile("file", file.Name())
	if err != nil {
		panic(err)
	}

	// Copy the file contents to the form field
	_, err = io.Copy(fileField, file)
	if err != nil {
		panic(err)
	}

	// Close the multipart writer
	writer.Close()

	// Send the HTTP request
	url := "https://apim.mushroomscan.com/api/predict"
	req, err := http.NewRequest("POST", url, requestBody)
	if err != nil {
		fmt.Println(err)
	}
	req.Header.Set("Content-Type", writer.FormDataContentType())

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	var response string
	err = json.Unmarshal(bodyBytes, &response)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println(response)
}
