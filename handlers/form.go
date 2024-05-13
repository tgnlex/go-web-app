package handlers

import (
	"fmt"
	"io"
	"net/http"
)

func FormHandle(w Res, r Req) {
	ctx := r.Context()
	val := r.PostFormValue("hi")

	fmt.Printf("%s: got /hello request \n", ctx.Value(val))

	if val == "" {
		w.Header().Set("x-missing-field", "val")
		w.WriteHeader(http.StatusBadRequest)
	}
	io.WriteString(w, fmt.Sprintf("Hello, %s!\n", val))
}
