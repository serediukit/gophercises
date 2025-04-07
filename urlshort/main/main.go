package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/boltdb/bolt"
	"github.com/serediukit/gophercises/urlshort"
)

var (
	pathsToUrls = map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}

	yaml = `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	json = `
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
)

var (
	yamlPath   = flag.String("yaml", "", "Path to YAML file")
	jsonPath   = flag.String("json", "", "Path to JSON file")
	boltDBPath = flag.String("bolt", "mybolt.db", "Path to Bolt DB file")
)

func init() {
	flag.Parse()
}

func main() {
	mux := defaultMux()

	var err error
	if *yamlPath != "" {
		yaml, err = readFileToString(*yamlPath)
		if err != nil {
			panic(err)
		}
	}
	if *jsonPath != "" {
		json, err = readFileToString(*jsonPath)
		if err != nil {
			panic(err)
		}
	}

	db, err := bolt.Open(*boltDBPath, 0600, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		err := db.Close()
		if err != nil {
			return
		}
	}()

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	jsonHandler, err := urlshort.JSONHandler([]byte(json), yamlHandler)
	if err != nil {
		panic(err)
	}

	boltDBHandler, err := urlshort.BoltDBHandler(db, jsonHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")

	err = http.ListenAndServe(":8080", boltDBHandler)
	if err != nil {
		panic(err)
	}
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, _ *http.Request) {
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
