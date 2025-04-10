package main

import (
	"fmt"
	"html-link-parser/utils"
	"os"
)

var path = "ex.html"

func main() {
	if len(os.Args) > 1 {
		path = os.Args[1]
	}

	res, err := utils.ReadFileToByte(path)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(res))
}
