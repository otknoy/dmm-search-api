package model

type SearchQuery struct {
	Keyword string
	Actress string
	Genre   string
	Maker   string
	Page    int
	N       int
}

type FacetQuery struct {
	Keyword string
	Actress string
	Genre   string
	Maker   string
}
