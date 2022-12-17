package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/config"
	"github.com/salamanderman234/peripheral-api/entity"
)

func Jwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		var claims entity.JWTClaims
		token, err := ctx.Cookie("token")
		if err != nil || token.Value == "" {
			return ctx.JSON(http.StatusUnauthorized, entity.BaseResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: "Missing cookie token",
			})
		}

		_, err = jwt.ParseWithClaims(token.Value, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.GetAppSecretKey()), nil
		})
		if err != nil {
			return ctx.JSON(http.StatusUnauthorized, entity.BaseResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: "Invalid token",
			})
		}

		return next(ctx)
	}
}
