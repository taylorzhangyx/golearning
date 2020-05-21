package main

import (
	"fmt"
	"net/http"
)

type zhang struct {
	location string
}

func (zhang) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, Zhang Family!")
}

func main() {
	var yuxin zhang
	http.ListenAndServe(":8080", yuxin)
}
