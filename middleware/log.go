package middleware

import (
	"github.com/labstack/echo/v4"
	utility "github.com/salamanderman234/peripheral-api/utility/log"
)

func Log(next echo.HandlerFunc) echo.HandlerFunc {
	return func(ctx echo.Context) error {
		utility.NewLogEntry(ctx).Info("incoming request")
		return next(ctx)
	}
}
