package router

import (
	"net/http"
	"strings"
	"toychart/handler"
	appMiddleware "toychart/middleware"
	"toychart/utils"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/mileusna/useragent"
	"gorm.io/gorm"
)

func SetupRoutes(h *handler.Handler, db *gorm.DB) *echo.Echo {
	e := echo.New()
	e.Validator = utils.NewValidator()

	e.Use(
		middleware.Recover(),
		middleware.Logger(),
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				agent := c.Request().UserAgent()
				ua := useragent.Parse(agent)
				if ua.Bot {
					return c.NoContent(http.StatusNoContent)
				}
				return next(c)
			}
		},
	)

	e.Use(
		func(next echo.HandlerFunc) echo.HandlerFunc {
			return func(c echo.Context) error {
				c.Response().Header().Set("X-Content-Type-Options", "nosniff")
				c.Response().Header().Set("Strict-Transport-Security", "max-age=31536000")
				c.Response().Header().Set("X-Frame-Options", "DENY")
				c.Response().Header().Set("Cache-control", "no-store")
				c.Response().Header().Set("Pragma", "no-store")
				return next(c)
			}
		},
		middleware.CORSWithConfig(middleware.CORSConfig{
			AllowCredentials: true,
			MaxAge:           86400,
			AllowMethods:     []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
			ExposeHeaders: []string{
				"x-request-id",
				"content-type",
				"vary",
			},
		}),
	)

	extractToken := func(c echo.Context) string {
		return strings.ReplaceAll(c.Request().Header.Get("Authorization"), "Bearer ", "")
	}

	user1 := e.Group("/user")
	user1.GET("", h.GetUser)
	user1.POST("", h.CreateUser)

	v := e.Group("/v1", appMiddleware.Authenticated(extractToken, h.VerifyToken))

	// User
	user := v.Group("/user")
	user.GET("", h.GetUser)
	// user.POST("", h.CreateUser)
	user.POST("/update/:id", h.UpdateUser)
	user.DELETE("/delete/:id", h.DeleteUser)

	// Toy
	toy := v.Group("/toy")
	toy.POST("/search", h.SearchToy)

	ebay := v.Group("/ebay")
	ebay.POST("/create", h.EbayCreate)

	return e
}
