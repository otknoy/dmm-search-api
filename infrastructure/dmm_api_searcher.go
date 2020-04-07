package infrastructure

import (
	"dmm-search-api/domain/model"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"
)

type DmmApiSearcher struct {
	ApiId       string
	AffiliateId string
}

func (s *DmmApiSearcher) Search(query model.Query) model.SearchResult {
	client := &http.Client{Timeout: time.Duration(3) * time.Second}

	uri := s.buildURL(string(query.Keyword))

	resp, err := client.Get(uri.String())
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	response := parse(resp)

	log.Println(response)

	return mapping(response)
}

func (s *DmmApiSearcher) buildURL(keyword string) url.URL {
	q := url.Values{}
	q.Add("api_id", s.ApiId)
	q.Add("affiliate_id", s.AffiliateId)
	q.Add("site", "FANZA")
	q.Add("service", "digital")
	q.Add("floor", "videoa")
	q.Add("hits", "20")
	q.Add("offset", "1")
	q.Add("keyword", keyword)
	q.Add("sort", "date")

	u := url.URL{}
	u.Scheme = "https"
	u.Host = "api.dmm.com"
	u.Path = "affiliate/v3/ItemList"
	u.RawQuery = q.Encode()

	return u
}

func parse(r *http.Response) Response {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatal(err)
	}

	var response Response
	if err := json.Unmarshal(body, &response); err != nil {
		log.Fatal(err)
	}

	return response
}

func mapping(response Response) model.SearchResult {
	var items []model.Item
	for _, item := range response.Result.Items {
		items = append(items, model.Item{
			ID:       item.ContentID,
			Title:    item.Title,
			URL:      item.AffiliateURL,
			ImageURL: item.ImageURL.Small,
		})
	}

	return model.SearchResult{
		ResultCount:   response.Result.ResultCount,
		TotalCount:    response.Result.TotalCount,
		FirstPosition: response.Result.FirstPosition,
		Items:         items,
	}
}
