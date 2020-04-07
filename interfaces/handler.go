package interfaces

import (
	"dmm-search-api/domain/model"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

type SearchHandler struct {
}

func (h *SearchHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	qs := r.URL.Query()

	keyword := model.Keyword(strings.Join(qs["keyword"], " "))
	log.Println(keyword)

	searchResult := model.SearchResult{} // h.SearchService.Search(model.Query{Keyword: keyword})
	// log.Println(searchResult)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(searchResult)
}
