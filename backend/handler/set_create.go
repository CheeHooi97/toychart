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
		for _, ips := range company.IP {
			for _, series := range ips.Series {
				if series.Items == nil {
					continue
				}
				for _, item := range series.Items {
					checkIPTypeExists, err := h.IPType.CheckIPTypeExists(company.IPType)
					if err != nil {
						return responseError(c, errcode.InternalServerError)
					}

					var ipType *model.IPType
					if !checkIPTypeExists {
						ipType = model.NewIPType()
						ipType.IPTypeName = company.IPType
						// photoUrl

						if err := h.IPType.Create(ipType); err != nil {
							return responseError(c, errcode.InternalServerError)
						}
					} else {
						ipType, err = h.IPType.GetByName(company.IPType)
						if err != nil {
							return responseError(c, errcode.InternalServerError)
						}
					}

					var ip *model.IP
					checkIPExists, err := h.IP.CheckIPExists(ips.IP)
					if err != nil {
						return responseError(c, errcode.InternalServerError)
					}

					if !checkIPExists {
						ip = model.NewIP()
						ip.IPTypeId = ipType.Id
						ip.IPName = ips.IP
						// photoUrl

						if err := h.IP.Create(ip); err != nil {
							return responseError(c, errcode.InternalServerError)
						}
					} else {
						ip, err = h.IP.GetByIPName(ips.IP)
						if err != nil {
							return responseError(c, errcode.InternalServerError)
						}
					}

					var ipSeries *model.IPSeries
					checkIPSeriesExists, err := h.IPSeries.CheckIPSeriesExists(series.Name)
					if err != nil {
						return responseError(c, errcode.InternalServerError)
					}

					if !checkIPSeriesExists {
						ipSeries = model.NewIPSeries()
						ipSeries.IPTypeId = ipType.Id
						ipSeries.IPId = ip.Id
						ipSeries.Series = series.Name
						// photoUrl

						if err := h.IPSeries.Create(ipSeries); err != nil {
							return responseError(c, errcode.InternalServerError)
						}
					} else {
						ipSeries, err = h.IPSeries.GetBySeries(series.Name)
						if err != nil {
							return responseError(c, errcode.InternalServerError)
						}
					}

					checkIPSeriesItemExists, err := h.IPSeriesItem.CheckIPSeriesItemExists(item.Name)
					if err != nil {
						return responseError(c, errcode.InternalServerError)
					}

					if !checkIPSeriesItemExists {
						ipItem := model.NewIPSeriesItem()
						ipItem.IPTypeId = ipType.Id
						ipItem.IPId = ip.Id
						ipItem.IPSeriesId = ipSeries.Id
						ipItem.ItemName = item.Name
						// photoUrl
						ipItem.Rarity = item.Rarity

						if err := h.IPSeriesItem.Create(ipItem); err != nil {
							return responseError(c, errcode.InternalServerError)
						}
					}
				}
			}
		}
	}

	return responseJSON(c, echo.Map{
		"success": true,
	})
}
