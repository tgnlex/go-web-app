package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"example.com/project/handlers"
	"github.com/joho/godotenv"
)

func main() {
	// Aliases
	print := fmt.Printf
	handle := http.HandleFunc
	timestamp := time.Now()
	listen := http.ListenAndServe
	warn := log.Fatal
	env := os.Getenv
	logline := log.Println
	// Load godotenv module
	err := godotenv.Load()
	if err != nil {
		logline("Error loading .env file")
	}
	// Get Enviornment Variables
	port := env("PORT")
	// Register Route Handlers
	handle("/", handlers.IndexHandle)
	handle("/submit", handlers.FormHandle)
	// Start Server
	print("[GO]: Starting http server at port %s \n", port)
	print("[GO]: Time {%s}", timestamp)
	if err := listen(port, nil); err != nil {
		warn(err)
	}

}
