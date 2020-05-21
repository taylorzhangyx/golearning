package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path { // exact string matching
	case "/":
		io.WriteString(w, "default")
	case "/dog":
		io.WriteString(w, "bart!")
	case "/cat":
		io.WriteString(w, "miao")
	}
}

func main() {
	var d hotdog
	http.ListenAndServe(":8080", d)
}
