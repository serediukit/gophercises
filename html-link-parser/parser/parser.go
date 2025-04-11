package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"html-link-parser/utils"
	"os"
)

func Parse(path string) {
	reader, err := utils.ReaderFromFile(path)
	if err != nil {
		fmt.Println("Can't open file")
		os.Exit(1)
	}
	z := html.NewTokenizer(reader)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		fmt.Println(z.Token().Data, "|", z.Token().Attr)
	}
}
