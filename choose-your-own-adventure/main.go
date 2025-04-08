package main

import (
	"bufio"
	"cyoa/decoder"
	"cyoa/entity"
	"cyoa/htmlPages"
	"flag"
	"fmt"
	"net/http"
	"os"
	"strconv"
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
		game(decoded)
	} else {
		startServer(decoded)
	}
}

func game(decoded map[string]*entity.Arc) {
	fmt.Println("Starting game ...")
	fmt.Println()

	scanner := bufio.NewScanner(os.Stdin)

	currentArc := "intro"

GAME:
	for {
		arc := decoded[currentArc]
		fmt.Println()
		fmt.Println(arc.Title)
		fmt.Println()
		for _, story := range arc.Story {
			fmt.Println(story)
		}
		fmt.Println()
		if len(arc.Options) > 0 {
			fmt.Println("What do you want to do?")
			fmt.Println()
			for i, option := range arc.Options {
				fmt.Printf("%d. %s\n", i+1, option.Text)
			}
			fmt.Printf("%d. Exit\n", len(arc.Options)+1)
			fmt.Println()
			fmt.Print("Your choice: ")

			var choice int
			for {
				scanner.Scan()
				input := scanner.Text()
				c, err := strconv.Atoi(input)
				if err != nil || c >= len(arc.Options)+1 {
					fmt.Printf("Invalid choice. Your input must be a number between 1 and %d.\n", len(arc.Options)+1)
				} else {
					choice = c
					break
				}
			}

			switch {
			case choice <= len(arc.Options):
				currentArc = arc.Options[choice-1].Arc
			default:
				break GAME
			}
		} else {
			fmt.Println("It's the END.")
			fmt.Println()
			break GAME
		}
	}
}

func startServer(decoded map[string]*entity.Arc) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", redirectToIntro)

	// test route for show parsed data
	mux.HandleFunc("/intro/json", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, decoded)
		if err != nil {
			panic(err)
		}
	})

	for arcName, arc := range decoded {
		mux.HandleFunc("/"+arcName, htmlPages.NewHandler(arc))
	}

	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func redirectToIntro(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/intro", http.StatusFound)
}

func readFile(filename string) []byte {
	data, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	return data
}
