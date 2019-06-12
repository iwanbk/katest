package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Welcome to my website!\n")
	})

	http.HandleFunc("/healthcheck", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "I'm healthy!!!\n")
	})

	port := ":8181"

	log.Printf("starting katest server at %v", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
