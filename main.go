package main

import (
	"log"
)

func main() {
	port := 8080

	// http.HandleFunc("/search", handler.Search)
	// http.HandleFunc("/facet", handler.Facet)

	// log.Fatal(http.ListenAndServe(":"+string(port), nil))
	log.Print(port)
}
