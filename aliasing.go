package main 

import(
	"fmt"
	"net/http"
	"os"
)
// Aliases
var Handle = http.HandleFunc
var Env = os.Getenv
var Print = fmt.Printf

// Lib
