package model

type SearchResponse struct {
	Query SearchQuery
	Total int
	Page  int
	N     int
	Items []Item
}

type FacetResponse struct {
	Query  FacetQuery
	Facets []Facet
}
