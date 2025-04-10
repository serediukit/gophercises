package main

import (
	"html-link-parser/parser"
	"os"
)

var path = "ex.html"

func main() {
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	parser.Parse(path)
}
