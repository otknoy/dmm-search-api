package service

import "github.com/otknoy/dmm-search-api/domain/model"

type ItemSearchService interface {
	SearchItems(model.SearchQuery) (model.SearchResponse, error)
}
