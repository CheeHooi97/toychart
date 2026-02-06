package handler

import (
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) EbaySearch(c echo.Context) error {
	var i struct {
		CompanyId string `json:"companyId"`
		Username  string `json:"username"`
		FcmToken  string `json:"fcmToken"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	return responseJSON(c, true)
}
