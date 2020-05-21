package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello, playground")
	http.HandleFunc("/", def)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}

func def(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "default")
}

func dog(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "doggy")
}

func me(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Yuxin Zhang")
}
