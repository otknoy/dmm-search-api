package main

import (
	"dmm-search-api/interfaces"
	"log"
	"net/http"
)

func main() {
	handler := &interfaces.SearchHandler{}

	mux := http.NewServeMux()
	mux.Handle("/search", handler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
