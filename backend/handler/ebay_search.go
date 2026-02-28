package handler

import (
	"fmt"
	"time"
	"toychart/errcode"
	"toychart/kit/ebay"
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) EbaySearch(c echo.Context) error {
	var i struct {
		Keyword string `json:"keyword"`
		Set     string `json:"set"`
		Order   string `json:"order"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	result, err := ebay.SearchToyPrices("Charizard")
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	fmt.Println("Ori Length: ", len(result.ItemSummaries))

	now := time.Now().UTC()

	soldItems := make([]struct {
		Title string
		Price string
		Url   string
	}, 0)

	for _, item := range result.ItemSummaries {
		// Must be auction
		isAuction := false
		for _, opt := range item.BuyingOptions {
			if opt == "AUCTION" {
				isAuction = true
				break
			}
		}
		if !isAuction {
			continue
		}

		// Must have ended
		if item.ItemEndDate == "" {
			continue
		}

		endTime, err := time.Parse(time.RFC3339, item.ItemEndDate)
		if err != nil {
			continue
		}

		if endTime.After(now) {
			continue // still bidding
		}

		// ✅ This is a truly SOLD item
		soldItems = append(soldItems, struct {
			Title string
			Price string
			Url   string
		}{
			Title: item.Title,
			Price: item.Price.Value,
			Url:   item.ItemWebUrl,
		})
	}

	fmt.Println("New Length: ", len(soldItems))

	return responseJSON(c, result)
}
