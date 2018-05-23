package model

type Item struct {
	ID       string   `json:"id"`
	Title    string   `json:"title"`
	URL      string   `json:"url"`
	ImageURL string   `json:"image_url"`
	Actress  []string `json:"actress"`
	Genre    []string `json:"genre"`
	Maker    []string `json:"maker"`
}
