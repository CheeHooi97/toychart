package handler

import (
	"net/http"
	"toychart/errcode"
	"toychart/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type Handler struct {
	Token        *service.TokenService
	Toy          *service.ToyService
	ToyPrice     *service.ToyPriceService
	IP           *service.IPService
	IPType       *service.IPTypeService
	IPSeries     *service.IPSeriesService
	IPSeriesItem *service.IPSeriesItemService
	User         *service.UserService
	UserDevice   *service.UserDeviceService
	DB           *gorm.DB
}

func NewHandler(services *service.Services, db *gorm.DB) *Handler {
	h := &Handler{

		Token:        services.TokenService,
		Toy:          services.ToyService,
		ToyPrice:     services.ToyPriceService,
		IP:           services.IPService,
		IPType:       services.IPTypeService,
		IPSeries:     services.IPSeriesService,
		IPSeriesItem: services.IPSeriesItemService,
		User:         services.UserService,
		UserDevice:   services.UserDeviceService,
		DB:           db,
	}

	return h
}

func responseError(c echo.Context, message errcode.ErrorCode) error {
	return c.JSON(http.StatusOK, map[string]any{
		"result": nil,
		"errmsg": message.Message,
		"error":  true,
		"status": false,
	})
}

func responseJSON(c echo.Context, result any) error {
	return c.JSON(http.StatusOK, map[string]any{
		"result": result,
		"errmsg": "",
		"error":  false,
		"status": true,
	})
}

func responseListJSON(c echo.Context, result any) error {
	return c.JSON(http.StatusOK, map[string]any{
		"result": map[string]any{
			"groups": result,
		},
		"errmsg": "",
		"error":  false,
		"status": true,
	})
}

func responseNoContent(c echo.Context) error {
	return c.JSON(http.StatusOK, nil)
}

func responseValidationError(c echo.Context, message string) error {
	return c.JSON(http.StatusOK, map[string]any{
		"result": nil,
		"errmsg": message,
		"error":  true,
		"status": false,
	})
}
