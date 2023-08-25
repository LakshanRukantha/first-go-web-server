package main

import (
	"fmt"
	"log"
	"net/http"
)


func main_handler(w http.ResponseWriter, r *http.Request){
	server_name := "GO Server"
	port := "8080"
	fmt.Fprintf(w, "<h1>🚀 %s is operational on Port: %s</h1>", server_name, port)
}

func hello(w http.ResponseWriter, r *http.Request){
	fmt.Fprint(w, "<h1>Hello World</h1>")
}

func main() {
	http.HandleFunc("/", main_handler)
	http.HandleFunc("/hello", hello)
	log.Fatal(http.ListenAndServe(":8080", nil))
}