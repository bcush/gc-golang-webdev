// * contains information about restaurant's menu including Breakfast, Lunch, and Dinner items

package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

type menu struct {
	breakfast meal
	lunch     meal
	dinner    meal
}

type item struct {
	Name string
}

type meal struct {
	Name  string
	Items []item
}

func main() {
	breakfast := meal{
		"Breakfast",
		[]item{
			struct {
				Name string
			}{
				"coffee",
			},

			struct {
				Name string
			}{
				"milk",
			},
		},
	}

	lunch := meal{
		"lunch",
		[]item{
			struct {
				Name string
			}{
				"tacos",
			},

			struct {
				Name string
			}{
				"beer",
			},

			struct {
				Name string
			}{
				"water",
			},
		},
	}

	dinner := meal{
		"dinner",
		[]item{
			struct {
				Name string
			}{
				"steak",
			},
		},
	}

	data1 := menu{breakfast, lunch, dinner}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", data1)
	if err != nil {
		log.Fatalln(err)
	}
}
