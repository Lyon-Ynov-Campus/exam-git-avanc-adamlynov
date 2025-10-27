package main

import (
	"fmt"
	"net/http"
)

func ColorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Colors</h1><style>*{background-color: #006400;}</style>")
}

func main() {
	http.HandleFunc("/color", ColorHandler)
	http.ListenAndServe(":8080", nil)
}
