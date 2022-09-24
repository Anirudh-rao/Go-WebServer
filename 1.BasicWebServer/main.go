package main

import (
	"fmt"
	"net/http"
)

const port = ":9090"

func Home(w http.ResponseWriter, r *http.Request) {
	response := "Hello World"
	fmt.Fprintf(w, response)
}

func main() {
	http.HandleFunc("/", Home)
	http.ListenAndServe(port, nil)
}
