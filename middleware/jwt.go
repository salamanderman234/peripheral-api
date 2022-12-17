package middleware

import (
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"github.com/labstack/echo/v4"
	"github.com/salamanderman234/peripheral-api/config"
	"github.com/salamanderman234/peripheral-api/entity"
	"github.com/salamanderman234/peripheral-api/utility"
)

func Jwt(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		token, err := ctx.Cookie("token")
		if err != nil || token.Value == "" {
			go utility.NewLogEntry(ctx).Error(err)
			return ctx.JSON(http.StatusUnauthorized, entity.BaseResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: "Missing cookie token",
			})
		}

		_, err = jwt.Parse(token.Value, func(t *jwt.Token) (interface{}, error) {
			return []byte(config.GetAppSecretKey()), nil
		})
		if err != nil {
			go utility.NewLogEntry(ctx).Error(err)
			return ctx.JSON(http.StatusUnauthorized, entity.BaseResponse{
				Code:   http.StatusUnauthorized,
				Status: "Unauthorized",
				Errors: "Invalid token",
			})
		}

		return next(ctx)
	}
}
