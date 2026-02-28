package handler

import (
	"toychart/config"
	"toychart/errcode"
	"toychart/middleware"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
)

func (h Handler) VerifyToken(c echo.Context, token string) (*middleware.Actor, error) {
	if token == "" {
		return nil, responseError(c, errcode.TokenNotFound)
	}

	tk, err := jwt.ParseWithClaims(token, &jwt.RegisteredClaims{}, func(t *jwt.Token) (interface{}, error) {
		return config.AuthenticationPublicKey, nil
	})

	if err != nil || !tk.Valid {
		return nil, responseError(c, errcode.InvalidToken)
	}

	claims, ok := tk.Claims.(*jwt.RegisteredClaims)
	if !ok || claims.ID == "" {
		return nil, responseError(c, errcode.InvalidToken)
	}

	actor := &middleware.Actor{
		Id:    claims.ID,
		Token: token,
	}

	return actor, nil
}
