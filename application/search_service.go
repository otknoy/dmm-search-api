package application

import (
	"dmm-search-api/domain"
	"dmm-search-api/domain/model"
)

type SearchService struct {
	ItemSearcher domain.ItemSearcher
}

func (s *SearchService) Search(query model.Query) model.SearchResult {
	return s.ItemSearcher.Search(query)
}
