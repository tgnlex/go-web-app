package utils

import (
	"log"

	"github.com/joho/godotenv"
)

// Aliases
// var handle = http.HandleFunc
// var print = fmt.Printf
// var warn = log.Fatal
var logger = log.Println
var env = godotenv.Load()
