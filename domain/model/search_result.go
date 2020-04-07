package model

type ResultCount int

type TotalCount int

type FirstPosition int

type SearchResult struct {
	ResultCount   ResultCount   `json:result_count`
	TotalCount    TotalCount    `json: total_count`
	FirstPosition FirstPosition `json: first_position`
	Items         []Item        `json: items`
}
