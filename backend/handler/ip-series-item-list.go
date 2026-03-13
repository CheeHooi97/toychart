package handler

import (
	"toychart/errcode"
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) IPSeriesItemList(c echo.Context) error {
	var i struct {
		IPSeriesId string `json:"ipSeriesId" validate:"required"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	ipItemsList, err := h.IPSeriesItem.GetByIPSeriesId(i.IPSeriesId)
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	return responseJSON(c, echo.Map{
		"lists": ipItemsList,
	})
}
