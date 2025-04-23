package ex5

import (
	"fmt"
	"net/http"
	"unicode/utf8"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	len := utf8.RuneCountInString(name)
	fmt.Fprintf(w, "Hello, %v, your name has %v letters", name, len)
}
