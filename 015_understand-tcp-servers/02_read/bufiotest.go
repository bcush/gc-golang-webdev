package main

import (
	"bufio"
	"fmt"
	"strings"
)

func main() {
	s := "This is a stringy string"

	scanner := bufio.NewScanner(strings.NewReader(s))
	scanner.Split(bufio.ScanRunes)

	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
	}

}
