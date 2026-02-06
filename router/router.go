package router

import (
	"toychart/handler"
	"toychart/middleware"
	"toychart/utils"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SetupRoutes(h *handler.Handler, db *gorm.DB) *echo.Echo {
	e := echo.New()
	e.Validator = utils.NewValidator()

	v := e.Group("/v1", middleware.Authenticate(db))

	// User
	user := v.Group("/user")
	user.GET("", h.GetUser)
	user.POST("", h.CreateUser)
	user.POST("/update/:id", h.UpdateUser)
	user.DELETE("/delete/:id", h.DeleteUser)

	// Card
	card := v.Group("/card")
	card.POST("/search", h.SearchCard)

	// set
	set := v.Group("/set")
	set.POST("/create", h.CreateSet)
	set.GET("/list", h.SetList)

	ebay := v.Group("/ebay")
	ebay.POST("/create", h.EbayCreate)

	return e
}
