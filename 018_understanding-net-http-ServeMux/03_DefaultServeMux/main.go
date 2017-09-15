package main

import (
	"io"
	"net/http"
)

type hotdog int
type hotcat int

func (m hotdog) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "HELLO DOGZZ")
}

func (m hotcat) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "HELLO CAKZT KATZ")
}

func main() {
	var d hotdog
	var c hotcat

	http.Handle("/cat", c)
	http.Handle("/cat/", c)
	http.Handle("/dog", d)
	http.Handle("/dog/", d)

	http.ListenAndServe(":8080", nil)
}
