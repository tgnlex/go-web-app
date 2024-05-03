package main

import (
	"fmt"
	"log"
	"net/http"
	"example.com/project/handlers"
)

func main() {
	http.HandleFunc("/hello", handlers.IndexHandler)

	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
