package main

import (
	"io"
	"net/http"
)

func d(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "heLLOOOD dog")
}

func c(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "hello CAT")
}

func main() {
	http.HandleFunc("/cat", c)
	http.HandleFunc("/dog", d)

	http.ListenAndServe(":8080", nil)
}
