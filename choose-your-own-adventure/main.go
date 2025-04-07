package main

import (
	"fmt"
	"net/http"
	"os"

	"cyoa/decoder"
	"cyoa/htmlPages"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)

	filepath := "res/gopher.json"
	data := readFile(filepath)
	decoded, err := decoder.Decode(data)
	if err != nil {
		panic(err)
	}
	mux.HandleFunc("/story", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, decoded)
		if err != nil {
			panic(err)
		}
	})

	for arcName, arc := range decoded {
		mux.HandleFunc("/"+arcName, htmlPages.NewHandler(arc))
	}

	fmt.Println("Starting server on port 8080")
	err = http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func index(w http.ResponseWriter, _ *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome to the home page!")
	if err != nil {
		panic(err)
	}
}

func readFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}
