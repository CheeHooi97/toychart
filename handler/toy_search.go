package handler

import (
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SearchCard(c echo.Context) error {
	var i struct {
		Keyword string `json:"keyword"`
		Set     string `json:"set"`
		Order   string `json:"order"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	return responseJSON(c, echo.Map{
		"lists": nil,
	})
}
