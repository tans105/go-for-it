package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", home)
	http.ListenAndServe(port, nil)
}

func home(w http.ResponseWriter, req *http.Request) {
	fmt.Fprint(w, "Welcome to the homepage")
}
