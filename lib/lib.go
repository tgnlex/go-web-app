package lib

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/joho/godotenv"
)

var Print = fmt.Printf
var Warn = log.Fatal
var Logline = log.Println

var timestamp = time.Now()

func Server(ADDR string) {
	Print("[GO]: Starting http server at: %v \n", ADDR)
	Print("[GO]: Timestamp: %v", timestamp)
	if err := http.ListenAndServe(ADDR, nil); err != nil {
		Warn(err)
	}
}

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		Logline("Error Loading .env file.")
	}
}
