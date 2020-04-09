package main

import (
	"dmm-search-api/application"
	"dmm-search-api/infrastructure"
	"dmm-search-api/interfaces"
	"log"
	"net/http"
)

func main() {
	handler := &interfaces.SearchHandler{
		SearchService: application.SearchService{
			ItemSearcher: &infrastructure.DmmApiSearcher{
				ApiId:       "Z5kd6ymEvRSNneLSP24Y",
				AffiliateId: "exsearch-990",
			},
		},
	}

	mux := http.NewServeMux()
	mux.Handle("/search", handler)

	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
