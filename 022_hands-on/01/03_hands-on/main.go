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
	http.HandleFunc("/", index)
	http.HandleFunc("/dog/", dog)
	http.HandleFunc("/me/", me)

	http.ListenAndServe(":8080", nil)
}
