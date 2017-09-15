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
	buddha := sage{
		Name:  "Buddha",
		Motto: "The belief of no beliefs",
	}

	mlk := sage{
		Name:  "MLK",
		Motto: "He had a dream",
	}

	gandhi := sage{
		Name:  "Ghandi",
		Motto: "Other stuff",
	}

	xs := []sage{buddha, mlk, gandhi}

	err := tpl.Execute(os.Stdout, xs)
	if err != nil {
		log.Fatalln(err)
	}

}
