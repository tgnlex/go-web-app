package handlers

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl = template.Must(template.ParseFiles("static/index.html"))

func IndexHandle(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, nil)
}

func FormHandle(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	fmt.Printf("%s: got /hello request \n"ctx.Value())

	val := r.PostFormValue("val")
	if val == "" {
		w.Header().Set("x-missing-field", "val")
		w.WriteHeader(http.StatusBadRequest)
	}
	io.WriteString(w, fmt.Sprintf("Hello, %s!\n", val))
}
