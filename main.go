package main

import (
	"fmt"
<<<<<<< HEAD
	"log"
	"net/http"
)

func URLHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>URL visitée</h1><p>Vous avez visité : %s</p>", r.URL.Path)
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "<h1>Site Web en Go</h1><p>Page d'accueil</p>")
	})

	http.HandleFunc("/url", URLHandler)

	fmt.Println("Serveur démarré sur http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
=======
	"net/http"
)

func ColorHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Colors</h1><style>*{background-color: #006400;}</style>")
}

func main() {
	http.HandleFunc("/color", ColorHandler)
	http.ListenAndServe(":8080", nil)
>>>>>>> feature-color
}
