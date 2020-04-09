package infrastructure

type Response struct {
	Request struct {
		Parameters struct {
			AffiliateID string `json:"affiliate_id"`
			APIID       string `json:"api_id"`
			Floor       string `json:"floor"`
			Hits        string `json:"hits"`
			Keyword     string `json:"keyword"`
			Offset      string `json:"offset"`
			Service     string `json:"service"`
			Site        string `json:"site"`
			Sort        string `json:"sort"`
		} `json:"parameters"`
	} `json:"request"`
	Result struct {
		Status        int `json:"status"`
		ResultCount   int `json:"result_count"`
		TotalCount    int `json:"total_count"`
		FirstPosition int `json:"first_position"`
		Items         []struct {
			ContentID    string `json:"content_id"`
			ProductID    string `json:"product_id"`
			Title        string `json:"title"`
			URL          string `json:"URL"`
			AffiliateURL string `json:"affiliateURL"`
			ImageURL     struct {
				List  string `json:"list"`
				Small string `json:"small"`
				Large string `json:"large"`
			} `json:"imageURL"`
			Prices struct {
				Price      string `json:"price"`
				Deliveries struct {
					Delivery []struct {
						Type  string `json:"type"`
						Price string `json:"price"`
					} `json:"delivery"`
				} `json:"deliveries"`
			} `json:"prices"`
			Date     string `json:"date"`
			Iteminfo struct {
				Genre []struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"genre"`
				Maker []struct {
					ID   int    `json:"id"`
					Name string `json:"name"`
				} `json:"maker"`
				// Actresses Actresses `json:"actress"`
			} `json:"iteminfo"`
		} `json:"items"`
	} `json:"result"`
}
