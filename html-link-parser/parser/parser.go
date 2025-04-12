package parser

import (
	"fmt"
	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
	"html-link-parser/utils"
	"os"
	"strings"
)

func Parse(path string) {
	reader, err := utils.ReaderFromFile(path)
	if err != nil {
		fmt.Println("Can't open file")
		os.Exit(1)
	}
	//z := html.NewTokenizer(reader)
	//for {
	//	tt := z.Next()
	//	if tt == html.ErrorToken {
	//		break
	//	}
	//	fmt.Println(strings.TrimSpace(z.Token().Data))
	//}
	node, _ := html.Parse(reader)
	resNodes := make([]linkNode, 0)
	findHrefDFS(&resNodes, node)
	for _, n := range resNodes {
		fmt.Printf("%#v\n", n)
	}
}

func findHrefDFS(resNodes *[]linkNode, node *html.Node) {
	for n := range node.ChildNodes() {
		if n.Type == html.ElementNode && n.DataAtom == atom.A {
			var href string
			for _, a := range n.Attr {
				if a.Key == "href" {
					href = a.Val
					break
				}
			}
			text := getAllText(n)
			link := linkNode{Href: href, Text: text}
			*resNodes = append(*resNodes, link)
		} else {
			findHrefDFS(resNodes, n)
		}
	}
}

func getAllText(node *html.Node) string {
	texts := make([]string, 0)
	for n := range node.ChildNodes() {
		if n.FirstChild == nil {
			texts = append(texts, n.Data)
		} else {
			texts = append(texts, getAllText(n))
		}
	}
	return strings.Join(texts, " ")
}
