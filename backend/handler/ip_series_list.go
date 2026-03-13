package handler

import (
	"toychart/errcode"
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) IPSeriesList(c echo.Context) error {
	var i struct {
		IpId string `json:"ipId" validate:"required"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	ipSeriesList, err := h.IPSeries.GetbyIpId(i.IpId)
	if err != nil {
		return responseError(c, errcode.InternalServerError)
	}

	return responseJSON(c, echo.Map{
		"lists": ipSeriesList,
	})
}
