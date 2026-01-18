package main

import (
	router "dondarrion91/unit-converter/internal/http"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "You are visiting %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)

	router.Router()

	fmt.Printf("Server listening on port %d\n", 8080)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
