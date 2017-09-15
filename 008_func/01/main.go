package main

import (
	"log"
	"os"
	"strings"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.New("").Funcs(fm).ParseGlob("*.gohtml"))
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

// Define funcmap
// uc is the strings method ToUpper
// ft is the function we created; firstThree
var fm = template.FuncMap{
	"uc": strings.ToUpper,
	"ft": firstThree,
}

// Define function; will be included in funcmap
func firstThree(s string) string {
	s = strings.TrimSpace(s)
	if len(s) >= 3 {
		s = s[:3]
	}
	return s
}

func main() {
	buddha := sage{
		Name:  "Buddha",
		Motto: "Love everything",
	}

	ghandi := sage{
		Name:  "Ghandi",
		Motto: "Also love everything",
	}

	ford := car{
		Manufacturer: "Ford",
		Model:        "pinto",
		Year:         1982,
	}

	mustang := car{
		Manufacturer: "GM",
		Model:        "big car",
		Year:         2000,
	}

	sages := []sage{buddha, ghandi}
	cars := []car{ford, mustang}

	data := struct {
		Sages []sage
		Cars  []car
	}{
		sages,
		cars,
	}

	err := tpl.Execute(os.Stdout, data)
	if err != nil {
		log.Fatalln(err)
	}
}
