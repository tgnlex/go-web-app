package handlers

import (
	"bytes"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

// Aliases
// var print = fmt.Printf
// var warn = log.Fatal
var logger = log.Println
var env = godotenv.Load()
var buf = &bytes.Buffer{}
var bufOut = buf.WriteTo
var serverError = http.StatusInternalServerError
