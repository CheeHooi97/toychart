package handler

import (
	"toychart/constant"
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) SetList(c echo.Context) error {
	var i struct {
		IP string `json:"ip" validate:"required"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	setList := constant.IPSet[i.IP]

	return responseJSON(c, echo.Map{
		"lists": setList,
	})
}
