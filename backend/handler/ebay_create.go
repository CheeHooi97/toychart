package handler

import (
	"toychart/utils"

	"github.com/labstack/echo/v4"
)

func (h *Handler) EbayCreate(c echo.Context) error {
	var i struct {
		// CompanyId string `json:"companyId"`
		// Username  string `json:"username"`
		// FcmToken  string `json:"fcmToken"`
	}

	if msg, err := utils.ValidateRequest(c, &i); err != nil {
		return responseValidationError(c, msg)
	}

	// result, err := ebay.SearchToyPrices("v^1.1#i^1#p^1#f^0#r^0#I^3#t^H4sIAAAAAAAA/+VYe2wURRjv9QWklCZixAA2x4KJirs7e3d7j03vwnGl9nj0daU8ojZzu7PXpXu7y+5c22tRmwawCjVRxJj+Y6GYaCQatRISg2hjhBglJArqX0BiUNQYIVGCGuvs9SjXSgDpEZt4/+zNN998832/7zUzoLd0zkM7a3deLnfMKhzqBb2FDgdXBuaUliyfV1S4sKQA5DA4hnqX9Rb3FX1fZcGkaghNyDJ0zULOrqSqWUKGGKRSpibo0FIsQYNJZAlYFGLhdWsFFwMEw9SxLuoq5YxWBylXQEJuL+ePu9ySX/ZwhKpdldmsBynEQQQ4yQ1kJEMARDJvWSkU1SwMNUzWA5eXBi4aeJo5XgAegecYj5/bTDlbkGkpukZYGECFMuoKmbVmjq43VhVaFjIxEUKFouGaWH04Wr2qrrmKzZEVyuIQwxCnrMmjiC4hZwtUU+jG21gZbiGWEkVkWRQbGt9hslAhfFWZ21A/A7XfI0Hgc8X9CHgln5/LC5Q1upmE+MZ62BRFouUMq4A0rOD0zRAlaMS3IBFnR3VERLTaaX8aU1BVZAWZQWrVyvCmcEMDFYq0IVSr6wqN9bTYBk1MNzRV03Lc4/Ih3u+lAz4eyD4gZjcal5aFecpOEV2TFBs0y1mn45WIaI0mY0OAycGGMNVr9WZYxrZGuXy+qxh6wWbbqeNeTOE2zfYrShIgnJnhzT0wsRpjU4mnMJqQMHUiA1GQgoahSNTUyUwsZsOnywpSbRgbAst2dnYynW5GNxOsCwCO3bhubUxsQ0lIEV4718f5lZsvoJWMKSIiKy1FwGmD6NJFYpUooCWoEM97fD4+i/tktUJTqf8g5NjMTs6IfGWI282T3JARj8g/UfbkI0NC2SBlbT1QHKbpJDTbETZUKCJaJHGWSiJTkQQ3L7vcfhnRkjcg056ALNNxXvLSnIwQQCgeFwP+/1Oi3Gqox5BoIpyXWM9bnG+trY6A+qZVTa7GjYmo3t1VkzAiHU1buwHP1oWBtKm6RcMS+4jRGA3eajZc1/iIqhBkmsn++QDAzvX8gVCrWxhJ0zIvJuoGatBVRUzPLAe7TamBZE86hlSVEKZlZNgwovmp1Xkz71+WiduzO3896j/qT9e1yrJDdmZZZa+3iABoKIzdgRhRT7J2ruuQHD9scmtGa+d1GacwsYRGGpaIGNKXpDgU2xkTQUnX1PS0cFPIyXdGoUbsHAdBkcaPrEwGCcbqEInFlp4iGFhMvX2Ca9bbkUb6ITZ1VUVmCzftepBMpjCMq2imFYY8JIgCZ1iz5nw+4OJ8nsD03CZmWnHrTCtpdikv7nPAO17OmxBUkzPLdsPUpZRon1HvwJWDnfwAEirI/Lg+xyjoc3xQ6HCAKnA/txQsKS1aX1w0d6GlYMQoUGYsJaGRe72JmHaUNqBiFs4v+PTUN3WV769+7ZlvF/TuWMa+UDAv5/1l6DFw78QLzJwiriznOQYsvjZTwlUsKHd5gQt4OB54eG4zWHpttpi7p/huuOKN/oPvNqz3ocSW9RVPrDnTmnodlE8wORwlBSRYCuT4T3cZ/V8c2v7lQvXE5Z6H7/uzMjb8zhXfe1d++2P1om1rDswqrex18aOHtdJfyI1gwcV+oXRo78CrbS0ReuyBgyOXTlaWj1RWjJ7a8OaOl7Xzbw9sD1/48YC+a2zMfWh42cjuisBgouvod0cWr6j7q6pnw+DF2R2B1melkud2nzvdHd701iunf9W2Pd99aWTP44tC4cNs49iHcNA4szwJv/48OnfxxwN8fPirxuPo0Y2Oxr3z6588e3yk/+yFeElHzWeBoz2//9y9L1JxLJhYYhpbHzzPle3pKyqTzy3ndlU/FRgcLlnRFIhu6zhZuovt2f/RS6OXT7749LHIvub9J+qWfDI7dOQHZu24L/8GuotNyRkTAAA=", "Labubu")
	// if err != nil {
	// 	return responseError(c, errcode.InternalServerError)
	// }

	return responseJSON(c, nil)
}
