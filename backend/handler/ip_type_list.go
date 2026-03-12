package handler

import (
	"toychart/constant"
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) IPTypeList(c echo.Context) error {
	var i struct {
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	ipList := constant.IPType

	return responseJSON(c, echo.Map{
		"lists": ipList,
	})
}
