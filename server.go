package main

import (
	"net/http"
	"time"
)

func Server(ADDR string) {
	Print("[GO]: Starting http server at: %v \n", ADDR)
	Print("[GO]: Timestamp: %v", time.Now())
	if err := http.ListenAndServe(ADDR, nil); err != nil {
		Warn(err)
	}
}
