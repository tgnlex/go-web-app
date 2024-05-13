package main

import (
	"fmt"
	"net/http"
	"os"

	"example.com/project/handlers"
	"example.com/project/utils"
)

var LoadEnv = utils.LoadEnv

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

