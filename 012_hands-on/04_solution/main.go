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

type hotel struct {
	Name, Address, City, Zip string
}

type region struct {
	Hotels []hotel
	Name   string
}

func main() {

	h1 := hotel{
		Name:    "Hard Rock",
		Address: "1234 West Elm",
		City:    "Los Angeles",
		Zip:     "912313",
	}

	h2 := hotel{
		Name:    "McDougals",
		Address: "9221 Pico",
		City:    "Los Angeles",
		Zip:     "91504",
	}

	h3 := hotel{
		Name:    "Donny Brucos",
		Address: "88812 Johnson St.",
		City:    "Whittier",
		Zip:     "123456",
	}

	h4 := hotel{
		Name:    "Applejacks",
		Address: "292929 Mountain",
		City:    "Pasadena",
		Zip:     "283922",
	}

	h5 := hotel{
		Name:    "Two Fish",
		Address: "923181 Argentine",
		City:    "Koreatown",
		Zip:     "923923",
	}

	h6 := hotel{
		Name:    "Sleep Zone",
		Address: "21333 Capital",
		City:    "Northridge",
		Zip:     "919922",
	}

	Southern := region{
		Hotels: []hotel{h1, h2},
		Name:   "Southern",
	}

	Central := region{
		Hotels: []hotel{h3, h4},
		Name:   "Central",
	}

	Northern := region{
		Hotels: []hotel{h5, h6},
		Name:   "Northern",
	}

	data := []region{Southern, Central, Northern}
	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}

}
