package handler

import (
	"strings"
	"toychart/errcode"
	"toychart/kit/ebay"
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) EbaySoldScrape(c echo.Context) error {
	var i struct {
		Keyword  string `json:"keyword"`
		URL      string `json:"url"`
		MaxPages int    `json:"maxPages"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	if strings.TrimSpace(i.Keyword) == "" && strings.TrimSpace(i.URL) == "" {
		return responseValidationError(c, "keyword or url is required")
	}

	result, err := ebay.ScrapeSoldToyPrices(strings.TrimSpace(i.Keyword), strings.TrimSpace(i.URL), i.MaxPages)
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	return responseJSON(c, result)
}
