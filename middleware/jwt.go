package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/entity"
)

func Jwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token, err := ctx.Cookie("token")
		if err != nil || token.Value == "" {
			return ctx.JSON(http.StatusUnauthorized, entity.BaseResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: "Missing cookie token",
			})
		}
	}
}
