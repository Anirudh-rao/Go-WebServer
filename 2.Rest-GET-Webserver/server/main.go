package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/Anirudh-rao/Go-WebServer/tree/main/2.Rest-GET-Webserver/countryCapitals"
)

const port = ":9000"

func home(w http.ResponseWriter, r *http.Request) {
	urlPaths := strings.Split(r.URL.Path, "/")
	if urlPaths[1] == "capital" {
		capital := countryCapitals.Capitals[urlPaths[2]]
		//If not match found, Empty string is returned
		if capital != "" {
			fmt.Fprintf(w, capital)
		} else {
			fmt.Fprintf(w, "sorry! No capital found For this Country")
		}
	}
}

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(port, nil)
}
