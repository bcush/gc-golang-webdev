package main

import (
	"log"
	"os"
	"strconv"
	"text/template"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("*.gohtml"))
}

func main() {
	arg1, _ := strconv.Atoi(os.Args[1])
	g1 := struct {
		Score1 int
		Score2 int
	}{
		arg1,
		9,
	}

	err := tpl.Execute(os.Stdout, g1)
	if err != nil {
		log.Fatalln(err)
	}
}
