package main

import (
	"io"
	"net/http"
)

type hotdog int

func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	io.WriteString(w, r.URL.Path)
	io.WriteString(w, "hello!")
}

type hotdf hotdog

func (d hotdf) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, r.URL.Path)
	io.WriteString(w, "Default")
}

func main() {
	var d hotdog
	var df hotdf
	mux := http.NewServeMux()
	mux.Handle("/", df)
	mux.Handle("/cat/", d)

	http.ListenAndServe(":8080", mux)
}
