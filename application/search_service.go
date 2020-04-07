package application

import (
	"dmm-search-api/domain/model"
)

type SearchService struct {
	// Searcher domain.Searcher
}

func (s *SearchService) Search(query model.Query) model.SearchResult {
	return model.SearchResult{}
}
