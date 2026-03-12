package handler

import (
	"toychart/errcode"
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SeriesList(c echo.Context) error {
	var i struct {
		Series string `json:"series" validate:"required"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	lists, err := h.Set.GetBySeries(i.Series)
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	return responseJSON(c, echo.Map{
		"lists": lists,
	})
}
