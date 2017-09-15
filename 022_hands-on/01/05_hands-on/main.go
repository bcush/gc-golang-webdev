package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func index(res http.ResponseWriter, req *http.Request) {
	d := "index"
	tpl.ExecuteTemplate(res, "index.gohtml", d)
}

func dog(res http.ResponseWriter, req *http.Request) {
	d := "dog"
	tpl.ExecuteTemplate(res, "index.gohtml", d)
}

func me(res http.ResponseWriter, req *http.Request) {
	d := "me"
	tpl.ExecuteTemplate(res, "index.gohtml", d)
}

func main() {
	http.Handle("/", http.HandlerFunc(index))
	http.Handle("/dog/", http.HandlerFunc(dog))
	http.Handle("/me/", http.HandlerFunc(me))

	http.ListenAndServe(":8080", nil)
}
