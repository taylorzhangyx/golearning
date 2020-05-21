package main

import (
	"html/template"
	"io"
	"log"
	"net/http"
)

type fooHan string
type barHan string
type mcleodHan string

func main() {
	var foo fooHan
	var bar barHan
	var mcleod mcleodHan
	http.Handle("/", foo)
	http.Handle("/dog/", bar)
	http.Handle("/me/", mcleod)
	http.ListenAndServe(":8080", nil)
}

func (fooHan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "foo ran")
}

func (barHan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bar ran")
}

func (mcleodHan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	tpl, err := template.ParseFiles("something.gohtml")
	if err != nil {
		log.Fatalln("error parsing template", err)
	}

	err = tpl.ExecuteTemplate(w, "something.gohtml", "McLeod")
	if err != nil {
		log.Fatalln("error executing template", err)
	}
}
