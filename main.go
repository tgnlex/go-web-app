package main

import (
	"fmt"
	"net/http"
	"os"

	"example.com/project/handlers"
	"example.com/project/lib"
)

// Aliases
var Handle = http.HandleFunc
var Env = os.Getenv
var Print = fmt.Printf

// Lib
var LoadEnv = lib.LoadEnv
var Server = lib.Server

// Handles
var Index = handlers.IndexHandle
var Form = handlers.FormHandle

func Main() {
	LoadEnv()
	// Get Enviornment Variables
	PORT := Env("PORT")
	// Register Route Handlers
	Handle("/", Index)
	Handle("/submit", Form)
	// Start Server
	Server(PORT)
}
