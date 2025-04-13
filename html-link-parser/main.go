package main

import (
	"fmt"
	"html-link-parser/parser"
	"os"
)

var path = "ex.html"

func main() {
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	links := parser.Parse(path)

	for _, n := range *links {
		fmt.Println(n)
	}
}
