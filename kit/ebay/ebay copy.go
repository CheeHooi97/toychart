package ebay

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"time"
)

type BrowseResponse struct {
	ItemSummaries []struct {
		Title string `json:"title"`
		Price struct {
			Value    string `json:"value"`
			Currency string `json:"currency"`
		} `json:"price"`
		ItemWebUrl string `json:"itemWebUrl"`
		Image      struct {
			ImageUrl string `json:"imageUrl"`
		} `json:"image"`
	} `json:"itemSummaries"`
}

func SearchToyPrices(token string, keyword string) (*BrowseResponse, error) {
	// Sandbox: https://api.sandbox.ebay.com/buy/browse/v1/item_summary/search
	endpoint := "https://api.ebay.com/buy/browse/v1/item_summary/search"

	params := url.Values{}
	params.Set("q", keyword)

	// 2. The Filter logic for Sold Items
	// In Browse API, we use the 'filter' parameter
	// conditions: 1000 is New, 3000 is Used.
	params.Set("filter", "lastItemsOnly:true,buyingOptions:{FIXED_PRICE}")
	params.Set("sort", "newlyListed")
	params.Set("limit", "10")

	reqURL := endpoint + "?" + params.Encode()

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	// 3. Set Headers (The Bearer token is mandatory)
	req.Header.Set("Authorization", "Bearer "+token)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("X-EBAY-C-MARKETPLACE-ID", "EBAY_US")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	var result BrowseResponse
	json.Unmarshal(body, &result)

	return &result, nil
}
