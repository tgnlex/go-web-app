package handlers

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("static/index.html"))
var buf = &bytes.Buffer{}
var BufOut = buf.WriteTo
var ServerError = http.StatusInternalServerError

type Res = http.ResponseWriter
type Req = *http.Request

func IndexHandle(w Res, r Req) {
	err := tpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), ServerError)
		return
	}
	BufOut(w)
}

func FormHandle(w Res, r Req) {
	ctx := r.Context()
	val := r.PostFormValue("val")

	fmt.Printf("%s: got /hello request \n", ctx.Value(val))

	if val == "" {
		w.Header().Set("x-missing-field", "val")
		w.WriteHeader(http.StatusBadRequest)
	}
	io.WriteString(w, fmt.Sprintf("Hello, %s!\n", val))
}
