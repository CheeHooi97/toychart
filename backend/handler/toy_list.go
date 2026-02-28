package handler

import (
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CardDetail(c echo.Context) error {
	var i struct {
		Id string `json:"id" validate:"required"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	return responseJSON(c, echo.Map{
		"detail": nil,
	})
}
