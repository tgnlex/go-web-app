package handlers

import (
	"net/http"
)

func saveHandler(w http.ResponseWriter, r *http.Request) {
	//	title, err := getTitle(w, r)
	//	if err != nil {
	//		return
	//	}
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	err := p.save()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	http.Redirect(w, r, "/view/"+title, http.StatusFound)
}
