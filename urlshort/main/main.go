package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"urlshort"
	// "github.com/serediukit/gophercises/urlshort"
)

var yamlPath = flag.String("yaml", "", "Path to YAML file")

func init() {
	flag.Parse()
}

func main() {
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	var yaml string
	var err error
	if *yamlPath != "" {
		yaml, err = readFileToString(*yamlPath)
		if err != nil {
			panic(err)
		}
	} else {
		yaml = `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	}

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	json := `
[
	{
		"path": "/serediuk",
		"url": "https://github.com/serediukit"
	},
	{	
		"path": "/serediuk-go",
		"url": "https://github.com/serediukit/gophercises"
	}
]
`

	jsonHandler, err := urlshort.JSONHandler([]byte(json), yamlHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")

	err = http.ListenAndServe(":8080", jsonHandler)
	if err != nil {
		panic(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintln(w, "Hello, world!")
	if err != nil {
		return
	}
}

func readFileToString(filePath string) (string, error) {
	bytes, err := os.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}
