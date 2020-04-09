package model

type SearchResult struct {
	ResultCount   int    `json:result_count`
	TotalCount    int    `json: total_count`
	FirstPosition int    `json: first_position`
	Items         []Item `json: items`
}
