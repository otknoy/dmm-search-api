package model

type Facet struct {
	Name  string
	Count int
	Items []FacetItem
}

type FacetItem struct {
	Value string
	Count int
}
