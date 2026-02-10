package ebay

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"
	"toychart/config"
)

type BrowseResponse struct {
	ItemSummaries []struct {
		ItemID     string `json:"itemId"`
		Title      string `json:"title"`
		ItemWebUrl string `json:"itemWebUrl"`

		BuyingOptions []string `json:"buyingOptions"`

		ItemEndDate string `json:"itemEndDate"`

		Price struct {
			Value    string `json:"value"`
			Currency string `json:"currency"`
		} `json:"price"`

		Image struct {
			ImageUrl string `json:"imageUrl"`
		} `json:"image"`
	} `json:"itemSummaries"`
}

func SearchToyPrices(keyword string) (*BrowseResponse, error) {
	// Sandbox: https://api.sandbox.ebay.com/buy/browse/v1/item_summary/search
	endpoint := "https://api.ebay.com/buy/browse/v1/item_summary/search"

	params := url.Values{}
	params.Set("q", keyword)

	params.Set("filter", "soldItemsOnly:true,buyingOptions:{AUCTION}")
	params.Set("sort", "-endDate")
	params.Set("limit", "200")

	// params.Set("filter", "lastItemsOnly:true,buyingOptions:{FIXED_PRICE}")
	// params.Set("sort", "newlyListed")
	// params.Set("limit", "10")

	reqURL := endpoint + "?" + params.Encode()

	req, err := http.NewRequest("GET", reqURL, nil)
	if err != nil {
		return nil, err
	}

	token, err := GetEbayAppToken(
		config.EbayClientId,
		config.EbayClientSecret,
	)
	if err != nil {
		log.Fatal(err)
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

func GetEbayAppToken(clientID, clientSecret string) (string, error) {
	data := url.Values{}
	data.Set("grant_type", "client_credentials")
	data.Set("scope", "https://api.ebay.com/oauth/api_scope")

	req, err := http.NewRequest(
		"POST",
		"https://api.ebay.com/identity/v1/oauth2/token",
		strings.NewReader(data.Encode()),
	)
	if err != nil {
		return "", err
	}

	// Basic Auth: base64(clientID:clientSecret)
	auth := base64.StdEncoding.EncodeToString([]byte(clientID + ":" + clientSecret))
	req.Header.Set("Authorization", "Basic "+auth)
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("token error: %s", body)
	}

	var result struct {
		AccessToken string `json:"access_token"`
		ExpiresIn   int    `json:"expires_in"`
	}

	if err := json.Unmarshal(body, &result); err != nil {
		return "", err
	}

	return result.AccessToken, nil
}
