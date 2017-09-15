package main

import (
	"io"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "HELLO INDEX")
}

func dog(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "HELLO DOG")
}

func me(res http.ResponseWriter, req *http.Request) {
	io.WriteString(res, "HELLO ME")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
