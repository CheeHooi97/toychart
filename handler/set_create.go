package handler

import (
	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateSet(c echo.Context) error {

	return responseJSON(c, echo.Map{
		"success": true,
	})
}
