package handler

import (
	"toychart/constant"
	"toychart/errcode"
	"toychart/model"

	"github.com/labstack/echo/v4"
)

func (h *Handler) CreateSet(c echo.Context) error {
	ipSeries := constant.IPSeries

	for _, company := range ipSeries {
		for _, ip := range company.IP {
			for _, series := range ip.Series {
				if series.Items == nil {
					continue
				}
				for _, item := range series.Items {
					set := model.NewSet()
					set.Series = series.Name
					// set.PhotoUrl = item.PhotoUrl
					set.ItemName = item.Name
					set.Rarity = item.Rarity
					set.IPType = company.IPType
					set.IPName = ip.IP
					// set.ItemTitle = item.Title
					// set.Href = item.Href

					if err := h.Set.Create(set); err != nil {
						return responseError(c, errcode.InternalServerError)
					}
				}
			}
		}
	}

	return responseJSON(c, echo.Map{
		"success": true,
	})
}
