package handler

import (
	"toychart/errcode"
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) IPList(c echo.Context) error {
	var i struct {
		IPTypeId string `json:"ipTypeId" validate:"required"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	ipList, err := h.IP.GetAllIPsByIPTypes(i.IPTypeId)
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	return responseJSON(c, echo.Map{
		"lists": ipList,
	})
}
