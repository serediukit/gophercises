package main

import (
	"cyoa/decoder"
	"flag"
	"os"

	"cyoa/engine"
)

var (
	console = flag.Bool("console", false, "Show output in console")
)

func init() {
	flag.Parse()
}

func main() {
	filepath := "res/gopher.json"
	data := readFile(filepath)
	decoded, err := decoder.Decode(data)
	if err != nil {
		panic(err)
	}

	if *console {
		engine.Game(decoded)
	} else {
		engine.StartServer(decoded)
	}
}

func readFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}
