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

type car struct {
	Manufacturer string
	Model        string
	Year         int
}

type data struct {
	Sages []sage
	Cars  []car
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

	ford := car{
		Manufacturer: "Ford",
		Model:        "Pinto",
		Year:         1982,
	}

	chevy := car{
		Manufacturer: "GM",
		Model:        "Truck",
		Year:         1999,
	}

	xsage := []sage{buddha, mlk, gandhi}
	xcar := []car{ford, chevy}

	results := data{
		Sages: xsage,
		Cars:  xcar,
	}

	err := tpl.Execute(os.Stdout, results)
	if err != nil {
		log.Fatalln(err)
	}

}
