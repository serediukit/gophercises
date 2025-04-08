package engine

import (
	"cyoa/entity"
	"cyoa/htmlPages"
	"fmt"
	"net/http"
)

func StartServer(decoded map[string]*entity.Arc) {
	mux := defaultMux(decoded)

	for arcName, arc := range decoded {
		mux.HandleFunc("/"+arcName, htmlPages.NewHandler(arc))
	}

	fmt.Println("Starting server on port 8080")
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		panic(err)
	}
}

func defaultMux(decoded map[string]*entity.Arc) *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", redirectToIntro)

	// test route for show parsed data
	mux.HandleFunc("/intro/json", showDecodedJSON(decoded))

	return mux
}

func redirectToIntro(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/intro", http.StatusFound)
}

func showDecodedJSON(decoded map[string]*entity.Arc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintln(w, decoded)
		if err != nil {
			panic(err)
		}
	}
}
