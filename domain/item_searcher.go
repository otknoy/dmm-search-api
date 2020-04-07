package domain

import "dmm-search-api/domain/model"

type ItemSearcher interface {
	Search(query model.Query) model.SearchResult
}
