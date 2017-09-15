package _1_stdout

import "fmt"

func main() {
	name := "Brian Cushman"

	tpl := `
	<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Hello World!</title>
	</head>
	<body>
	<h1>` + name + `</h1>
	</body>
	`

	fmt.Println(tpl)
}
