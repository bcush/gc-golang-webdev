package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))

}

type sage struct {
	Name  string
	Motto string
}

func main() {
	sg := sage{
		Name:  "Buddha",
		Motto: "All is well",
	}

	err := tpl.Execute(os.Stdout, sg)
	if err != nil {
		log.Fatalln(err)
	}
}
