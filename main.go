package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// Handler for the "/api" endpoint
func api_handler(w http.ResponseWriter, r *http.Request) {
	// Set the Content-Type header to indicate JSON response
	w.Header().Set("Content-Type", "application/json")

	// Define the API URL
	apiUrl := "https://random-data-api.com/api/v2/users"

	// Make an HTTP GET request to the API
	response, err := http.Get(apiUrl)
	if err != nil {
		// If an error occurs, respond with an error message
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer response.Body.Close()

	// Read the response body
	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the formatted and indented JSON response data
	var formattedJSON bytes.Buffer
	err = json.Indent(&formattedJSON, responseData, "", "  ")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Write the response data to the API response
	w.Write(formattedJSON.Bytes())
}

// Handler for the root ("/") endpoint
func main_handler(w http.ResponseWriter, r *http.Request) {
	server_name := "GO Server"
	port := "8080"
	fmt.Fprintf(w, "<h1>🚀 %s is operational on Port: %s</h1>", server_name, port)
}

// Handler for the "/hello" endpoint
func hello(w http.ResponseWriter, r *http.Request) {
	// Write a simple HTML response
	w.Write([]byte("<h1>Hello World</h1>"))
}

func main() {
	// Define handlers for different endpoints
	http.HandleFunc("/", main_handler)
	http.HandleFunc("/hello", hello)
	http.HandleFunc("/api", api_handler)

	// Start the HTTP server and listen on port 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}