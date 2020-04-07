package domain

import "dmm-search-api/domain/model"

type SearchService interface {
	Search(query model.Query) model.SearchResult
}

type Searcher interface {
	Search(query model.Query) model.SearchResult
}
