package repository

import "dmm-search-api/domain/model"

type ItemRepository interface {
	Search(query model.Query) *model.SearchResult
}
