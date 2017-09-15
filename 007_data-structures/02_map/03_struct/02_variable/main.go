package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sages struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	buddha := sages{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	err := tpl.Execute(os.Stdout, buddha)
	if err != nil {
		log.Fatalln(err)
	}

}
