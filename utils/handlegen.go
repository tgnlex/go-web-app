package utils

import (
	"net/http"
	"regexp"
)

var validPath = regexp.MustCompile("^/(index|blog|login)/([a-zA-Z0-9]+)$")

func MakeReqHandler(fn func(http.ResponseWriter, *http.Request, string)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		m := validPath.FindStringSubmatch(r.URL.Path)
		if m == nil {
			http.NotFound(w, r)
			return
		}
		fn(w, r, m[2])
	}
}
