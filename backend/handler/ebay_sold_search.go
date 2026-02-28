package handler

import (
	"toychart/errcode"
	"toychart/kit/ebay"
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) EbaySoldSearch(c echo.Context) error {
	var i struct {
		Keyword string `json:"keyword" validate:"required"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	result, err := ebay.SearchSoldToyPrices(i.Keyword)
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	return responseJSON(c, result)
}
