package infrastructure

import (
	"dmm-search-api/domain/model"
	"log"
)

type DmmApiSearcher struct {
}

func (s *DmmApiSearcher) Search(query model.Query) model.SearchResult {
	log.Println("hello")
	return model.SearchResult{}
}
