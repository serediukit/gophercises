package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	"github.com/serediukit/gophercises/urlshort"
)

var (
	yamlPath = flag.String("yaml", "", "Path to YAML file")
	jsonPath = flag.String("json", "", "Path to JSON file")
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

	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

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
